const {net, ipcMain, dialog} = require('electron')
const fs = require('fs')
const path = require('path')
const homedir = require('os').homedir()
const {spawn} = require('child_process')
const extract = require('extract-zip')

const isDev = () => {
  return process.env.NODE_ENV === 'development'
}

const netRequest = (url, cb) => {
  const request = net.request(url)
  request.on('response', (response) => {
    response.on('data', (chunk) => {
      cb(JSON.parse(chunk.toString()))
    })
  })
  request.end()
}

const getBND2Version = () => {
  return require('../package.json').version
}

const startKernel = (event) => {
  let fileName = 'bnd2.exe'
  if (process.platform !== 'win32') {
    fileName = 'bnd2'
  }

  let bnd2 = spawn(path.join(homedir, '.bnd2', fileName))

  event.sender.send('asynchronous-reply', {
    type: 'loaded',
  })
  const showTip = (data) => {
    event.sender.send('asynchronous-reply', {
      type: 'loaded',
    })
    event.sender.send('asynchronous-reply', {
      type: 'kernelError',
      data: data,
    })
  }
  bnd2.stdout.on('data', (data) => {
    showTip(data)
  })

  bnd2.stderr.on('data', (data) => {
    showTip(data)
  })

  bnd2.on('close', (code) => {
    showTip(`child process exited with code ${code}`)
  })
}

const cpFilename = (data) => {
  const aria2cFilename = process.platform === 'win32'
    ? 'aria2c_windows.zip'
    : 'aria2c_darwin.zip'
  extract(path.join(path.dirname(__dirname), aria2cFilename),
    {dir: `${homedir}/.bnd2`},
    () => {
      // chmod
      if (process.platform !== 'win32') {
        const aria2cfd = fs.openSync(homedir + '/.bnd2/aria2c', 'a')
        fs.fchmodSync(aria2cfd, '0777')
      }
    },
  )

  // cp bnd2
  const bnd2Filename = process.platform === 'win32' ? 'bnd2.exe' : 'bnd2'
  fs.writeFileSync(`${homedir}/.bnd2/${bnd2Filename}`,
    fs.readFileSync(path.join(path.dirname(__dirname), bnd2Filename)))
  if (process.platform !== 'win32') {
    const bnd2fd = fs.openSync(homedir + '/.bnd2/bnd2', 'a')
    fs.fchmodSync(bnd2fd, '0777')
  }

  // write .bnd2/KERNEL_VER
  fs.writeFileSync(homedir + '/.bnd2/KERNEL_VER', data.kernelVer, 'UTF-8')
}

const downloadKernel = (data, event) => {
  cpFilename(data)
  startKernel(event)
}

module.exports.isDev = isDev

module.exports.getBND2Version = getBND2Version

module.exports.netRequest = netRequest

module.exports.initIpcMain = () => {
  ipcMain.on('asynchronous-message', (event, arg) => {
    switch (arg) {
      case 'chooseFile' :
        dialog.showOpenDialog({
          defaultPath: homedir,
          properties: ['openDirectory', 'createDirectory'],
        }, (files) => {
          if (files) {
            event.sender.send('asynchronous-reply', {
              type: 'chooseFile',
              data: files,
            })
          }
        })
        break
      case 'checkVersion':
        netRequest('https://rhythm.b3log.org/version/bnd2', (data) => {
          // check version
          if (getBND2Version() < data.ver) {
            event.sender.send('asynchronous-reply', {
              type: arg,
              data: data,
            })
            return
          } else {
            event.sender.send('asynchronous-reply', {
              type: 'checkLogin',
            })
          }

          // check kernel version
          try {
            fs.statSync(homedir + '/.bnd2')
            try {
              const kernelVersion = fs.readFileSync(path.join(homedir,
                '/.bnd2/KERNEL_VER'), 'UTF-8')
              if (kernelVersion < data.kernelVer) {
                downloadKernel(data, event)
              } else {
                try {
                  fs.statSync(path.join(homedir,
                    (process.platform !== 'win32'
                      ? '/.bnd2/bnd2'
                      : '/.bnd2/bnd2.exe')))
                  startKernel(event)
                } catch (e) {
                  downloadKernel(data, event)
                }
              }
            } catch (e) {
              downloadKernel(data, event)
            }
          } catch (e) {
            fs.mkdirSync(`${homedir}/.bnd2`)
            downloadKernel(data, event)
          }
        })
        break
    }
  })
}