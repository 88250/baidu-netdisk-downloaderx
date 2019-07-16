import React from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableCell from '@material-ui/core/TableCell'
import TableRow from '@material-ui/core/TableRow'
import TableHead from '@material-ui/core/TableHead'
import TableSortLabel from '@material-ui/core/TableSortLabel'
import Tooltip from '@material-ui/core/Tooltip'
import Card from '@material-ui/core/Card'
import Button from '@material-ui/core/Button'
import Dialog from '@material-ui/core/Dialog'
import DialogActions from '@material-ui/core/DialogActions'
import DialogContent from '@material-ui/core/DialogContent'
import DialogContentText from '@material-ui/core/DialogContentText'
import DialogTitle from '@material-ui/core/DialogTitle'
import SaveAltIcon from '@material-ui/icons/SaveAlt'
import FolderIcon from '@material-ui/icons/Folder'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import { clearCookie } from '../utils/clearCookie'

export default class Index extends React.Component {

  state = {
    sort: {
      time: {
        active: false,
        direction: 'desc',
      },
      name: {
        active: true,
        direction: 'asc',
      },
    },
    list: [],
    breadcrumb: [],
    showPath: false,
    showTraverse: false,
    downloadRow: {},
    saveFilePath: localStorage.saveFilePath || '',
    traverseFinished: false,
    traverseData: {
      all: 1, dirs: 0, files: 1, hSize: '0 kB', finished: false,
    },
  }

  componentDidMount () {
    this.getList()
    window.ipcRenderer.on('asynchronous-reply', (event, arg) => {
      switch (arg.type) {
        case 'chooseFile':
          localStorage.saveFilePath = arg.data[0]
          this.setState({
            saveFilePath: arg.data[0],
          })
          break
        default:
          break
      }
    })
  }

  componentWillReceiveProps (nextProps, nextContent) {
    const data = nextProps.rwsData
    if (data.cmd === 'traverse') {
      if (!this.state.traverseFinished) {
        const traverseData = data.data
        this.setState({
          traverseFinished: traverseData.finished,
          traverseData,
        })
      }
    }
    if (data.cmd === 'ls') {
      this.setState({
        list: data.data || [],
        breadcrumb: this.breadcrumb(),
      })
    }
  }

  showPathDialog (downloadRow) {
    this.setState({
      showPath: true,
      downloadRow,
    })
  }

  closePathDialog () {
    this.setState({showPath: false})
  }

  choosePath () {
    window.ipcRenderer.send('asynchronous-message', 'chooseFile')
  }

  traverse () {
    this.setState({
      showPath: false,
    })

    if (this.state.downloadRow.isdir !== 1) {
      this.startDownload('downloadfile')
      return
    }

    this.setState({
      showTraverse: true,
      traverseFinished: false,
    })

    window.rws.send(JSON.stringify(
      {cmd: 'traverse', param: {path: this.state.downloadRow.path}}))
  }

  cancelTraverse () {
    this.setState({showTraverse: false})
    window.rws.send(JSON.stringify({cmd: 'canceltraverse', param: {}}))
  }

  startDownload (cmd) {
    this.setState({
      showTraverse: false,
    })

    window.rws.send(JSON.stringify({
      cmd,
      param: {
        path: this.state.downloadRow.path,
        saveDir: this.state.saveFilePath,
        size: this.state.downloadRow.size,
      },
    }))
  }

  getList (data, type) {
    if (!localStorage.currentPath) {
      localStorage.currentPath = '/'
    }

    if (data) {
      if (data.isdir !== 1 || data.path === '') {
        return
      }
      localStorage.currentPath = data.path
    }

    let order
    if (type) {
      let sort
      if (type === 'time') {
        order = this.state.sort.time.direction === 'desc' ? 'asc' : 'desc'
        sort = {
          time: {
            active: true,
            direction: order,
          },
          name: {
            active: false,
            direction: this.state.sort.name.direction,
          },
        }
      } else {
        order = this.state.sort.name.direction === 'asc' ? 'desc' : 'asc'
        sort = {
          time: {
            active: false,
            direction: this.state.sort.time.direction,
          },
          name: {
            active: true,
            direction: order,
          },
        }
      }
      this.setState({
        sort,
      })
    } else {
      if (this.state.sort.time.active) {
        type = 'time'
        order = this.state.sort.time.direction
      } else {
        type = 'name'
        order = this.state.sort.name.direction
      }
    }

    window.rws.send(
      JSON.stringify({
        cmd: 'ls',
        param: {
          path: localStorage.currentPath,
          by: type,
          order,
        },
      }),
    )
  }

  breadcrumb () {
    let breadcrumb = []
    if (localStorage.currentPath === '/') {
      return []
    }
    const pathList = localStorage.currentPath.split('/')
    pathList.forEach((name, index) => {
      if (index !== 0) {
        let path = ''
        pathList.forEach((data, i) => {
          if (i !== 0 && i <= index) {
            path += '/' + data
          }
        })

        if (index === pathList.length - 1) {
          path = ''
        }

        breadcrumb.push({
          name,
          path,
        })
      }
    })

    return breadcrumb
  }

  render () {
    const {classes} = this.props

    return (
      <div>
        <AppBar
          className={classes.menu}>
          <Toolbar>
            <Typography className={classes.fnFlex1} color="inherit" noWrap>
              全部文件
            </Typography>
            <Button
              className={classes.ftOriginal}
              color="inherit"
              onClick={clearCookie}
            >
              切换账号
            </Button>
          </Toolbar>
        </AppBar>
        <div className={classnames(classes.fnFlex, classes.ftNormalSize)}>
          <div className={classes.fnFlex1}>
            <span className={classes.link}
                  onClick={this.getList.bind(this,
                    {isdir: 1, path: '/'}, null)}>全部文件</span>
            {
              this.state.breadcrumb.map(row => {
                return (<span key={row.path}>
                  <span className={classes.ftGray}>&nbsp; > &nbsp;</span>
                  <span
                    className={row.path === '' ? classes.ftGray : classes.link}
                    onClick={this.getList.bind(this,
                      {isdir: 1, path: row.path}, null)}>{row.name}</span>
              </span>)
              })
            }
          </div>
          <div className={classes.ftGray}>
            共 {this.state.list.length} 个
          </div>
        </div>
        <br/>
        <Card>
          <Table className={classes.ftOriginal}>
            <TableHead>
              <TableRow>
                <TableCell width="50"></TableCell>
                <TableCell padding="none">
                  <Tooltip title="文件名排序">
                    <TableSortLabel
                      onClick={this.getList.bind(this, null, 'name')}
                      active={this.state.sort.name.active}
                      direction={this.state.sort.name.direction}>
                      文件名
                    </TableSortLabel>
                  </Tooltip>
                </TableCell>
                <TableCell width="180">
                  大小
                </TableCell>
                <TableCell padding="none" width="160">
                  <Tooltip title="修改时间排序">
                    <TableSortLabel
                      onClick={this.getList.bind(this, null, 'time')}
                      active={this.state.sort.time.active}
                      direction={this.state.sort.time.direction}>
                      修改时间
                    </TableSortLabel>
                  </Tooltip>
                </TableCell>
                <TableCell width="80">
                  操作
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {
                this.state.list.map(row => {
                  const isSelected = false
                  return (
                    <TableRow
                      hover
                      role="checkbox"
                      aria-checked={isSelected}
                      key={row.fs_id}
                      selected={isSelected}
                    >
                      <TableCell style={{paddingRight: 10 + 'px'}}>
                        {row.isdir === 1 ? <FolderIcon
                          className={classes.ftYellow}/> : ''}
                      </TableCell>
                      <TableCell
                        className={row.isdir === 1 ? classes.link : ''}
                        padding="none"
                        onClick={this.getList.bind(this, row, null)}>
                        {row.server_filename}
                      </TableCell>
                      <TableCell className={classes.ftGray}>
                        {row.isdir === 0 ? row.hSize : ''}
                      </TableCell>
                      <TableCell
                        padding="none" className={classes.ftGray}>
                        {row.hMtime}
                      </TableCell>
                      <TableCell>
                        <SaveAltIcon
                          onClick={this.showPathDialog.bind(this, row)}
                          className={classes.fnPointer}
                          color="action"/>
                      </TableCell>
                    </TableRow>
                  )
                })
              }
            </TableBody>
          </Table>
        </Card>
        <Dialog
          open={this.state.showTraverse}
          onClose={this.cancelTraverse.bind(this)}
        >
          <DialogTitle>文件夹信息</DialogTitle>
          <DialogContent>
            <DialogContentText className={classes.dialog}>
              下载名称： {this.state.downloadRow.server_filename} <br/>
              目录个数：{this.state.traverseData.dirs} <br/>
              文件个数：{this.state.traverseData.files} <br/>
              下载总计： {this.state.traverseData.all} <br/>
              下载大小：{this.state.traverseData.hSize}
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button
              className={classes.ftOriginal}
              color="primary"
              variant="contained"
              disabled={!this.state.traverseFinished}
              onClick={this.startDownload.bind(this, 'downloaddir')}>
              下载
            </Button>
            <Button
              className={classes.ftOriginal}
              variant="contained"
              onClick={this.cancelTraverse.bind(this)}>
              取消
            </Button>
          </DialogActions>
        </Dialog>
        <Dialog
          open={this.state.showPath}
          onClose={this.closePathDialog.bind(this)}
        >
          <DialogTitle>下载信息</DialogTitle>
          <DialogContent>
            <DialogContentText className={classes.dialog}>
              下载文件： {this.state.downloadRow.server_filename} <br/>
              下载目录：{this.state.saveFilePath}
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <div className={classes.fnFlex1}>
              <Button
                className={classes.ftOriginal}
                color="secondary"
                variant="contained"
                onClick={this.choosePath}>
                选择目录
              </Button>
            </div>
            <Button
              className={classes.ftOriginal}
              color="primary"
              variant="contained"
              onClick={this.traverse.bind(this)}>
              确定
            </Button>
            <Button
              className={classes.ftOriginal}
              variant="contained"
              onClick={this.closePathDialog.bind(this)}>
              取消
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    )
  }
}

Index.propTypes = {
  classes: PropTypes.object.isRequired,
}