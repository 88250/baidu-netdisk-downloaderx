const {ipcRenderer} = require('electron')
const {isDev, getBND2Version} = require('./src/utils')

const loginWV = document.querySelector('#loginWV')

const checkLogin = () => {
  let src = 'http://bnd2.b3log.org?' + (new Date()).getTime()
  const baiduURL = 'https://pan.baidu.com'
  const bnd2Cookies = {hp: '', bd: ''}

  if (isDev()) {
    loginWV.addEventListener('dom-ready', () => {
      loginWV.openDevTools()
    })
    src = 'http://localhost:3000'
  }

  loginWV.src = baiduURL
  loginWV.addEventListener('load-commit', () => {
    if (loginWV.src.indexOf(baiduURL) === 0) {
      loginWV.executeJavaScript('window.scrollTo(10000, 0)')
    }
    const session = loginWV.getWebContents().session

    session.cookies.get({url: baiduURL}, async (error, cookies) => {
      for (let i = 0; i < cookies.length; i++) {
        let cookie = cookies[i]
        if (cookie.name === 'BDUSS' && loginWV.src.indexOf(baiduURL) > -1) {
          bnd2Cookies.bd = cookie.value
          document.querySelector('#loginBaiduTip').style.display = 'none'
          await fetch('http://localhost:6804/login', {
            method: 'POST',
            body: JSON.stringify(bnd2Cookies),
          })
          loginWV.src = src
        }
      }

      if (bnd2Cookies.bd === '' && loginWV.src.indexOf(baiduURL) > -1) {
        document.querySelector('#loginBaiduTip').style.display = 'block'
      }
    })
  })
}

onload = () => {
  document.querySelector('title').innerHTML = 'BND2 v' + getBND2Version()
  ipcRenderer.on('asynchronous-reply', (event, arg) => {
    switch (arg.type) {
      case 'checkLogin':
        checkLogin()
        break
      case 'checkVersion':
        const downloadTip = document.querySelector('#downloadTip')
        downloadTip.onclick = () => {
          require('electron').
            shell.
            openExternal(arg.data.dl)
        }
        downloadTip.style.display = 'block'
        break
      case 'loaded':
        document.querySelector('#loadingTip').style.display = 'none'
        break
      case 'kernelError':
        if (!arg.data) {
          return
        }
        const errorTip = document.querySelector('#errorTip')
        errorTip.innerHTML = errorTip.innerHTML + arg.data + '<br>'
        errorTip.style.display = 'block'
        setTimeout(() => {
          errorTip.innerHTML = ''
          errorTip.style.display = 'none'
        }, 10000)
        break
    }
  })

  loginWV.addEventListener('ipc-message', (event) => {
    if (event.channel === 'clearCookie') {
      loginWV.getWebContents().session.clearStorageData()
      loginWV.src = baiduURL
      document.querySelector('#loginBaiduTip').style.display = 'block'
    }
  })

  ipcRenderer.send('asynchronous-message', 'checkVersion')
}