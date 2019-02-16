import React from 'react'
import PropTypes from 'prop-types'
import Button from '@material-ui/core/Button'
import LinearProgress from '@material-ui/core/LinearProgress'
import PlayIcon from '@material-ui/icons/PlayArrow'
import StopIcon from '@material-ui/icons/Stop'
import RemoveIcon from '@material-ui/icons/DeleteOutline'
import LocalIcon from '@material-ui/icons/FilterCenterFocus'

export default class Downloading extends React.Component {
  state = {
    list: [],
  }

  componentDidMount () {
    window.rws.send(JSON.stringify({cmd: 'lstasks', param: {}}))
  }

  componentWillReceiveProps (nextProps, nextContent) {
    const data = nextProps.rwsData
    if (data.cmd === 'lstasks') {
      this.setState({
        list: data.data,
      })
    }
  }

  componentWillUnmount () {
    window.rws.send(JSON.stringify({cmd: 'stoplstasks', param: {}}))
  }

  startAll () {
    window.rws.send(JSON.stringify({cmd: 'unpauseall', param: {}}))
  }

  stopAll () {
    window.rws.send(JSON.stringify({cmd: 'pauseall', param: {}}))
  }

  deleteAll () {
    window.rws.send(JSON.stringify({cmd: 'deldownloadall', param: {}}))
  }

  start (gid) {
    window.rws.send(JSON.stringify({cmd: 'unpause', param: {gid}}))
  }

  stop (gid) {
    window.rws.send(JSON.stringify({cmd: 'pause', param: {gid}}))
  }

  delete (gid) {
    window.rws.send(JSON.stringify({cmd: 'deldownload', param: {gid}}))
  }

  local (saveDir, name) {
    window.shell.showItemInFolder(window.path.join(saveDir, name))
  }

  render () {
    const {classes} = this.props

    return (
      <div>
        <div className={classes.fnFlex}>
          <div className={classes.fnFlex1}>
            <Button className={classes.ftOriginal}
                    color='primary'
                    onClick={this.startAll}
                    variant="contained">
              <PlayIcon/>
              全部开始
            </Button> &nbsp; &nbsp;
            <Button className={classes.ftOriginal}
                    onClick={this.stopAll}
                    variant="contained">
              <StopIcon/>
              全部暂停
            </Button> &nbsp; &nbsp;
            <Button className={classes.ftOriginal}
                    onClick={this.deleteAll}
                    color='secondary' variant="contained">
              <RemoveIcon/>
              全部删除
            </Button>
          </div>
          <div className={classes.ftGray}>
            共 {this.state.list.length} 个
          </div>
        </div>
        <br/>
        {
          this.state.list.map(row => {
            return (<div className={classes.listItem} key={row.gid}>
              <div className={classes.fnFlex}>
                <div className={classes.listTitle}>{row.name}</div>
                <div>
                  {row.state === 3 && (
                    <PlayIcon className={classes.fnPointer}
                              onClick={this.start.bind(this, row.gid)}
                              color='primary'/>)}

                  {row.state === 0 && (
                    <StopIcon className={classes.fnPointer}
                              onClick={this.stop.bind(this, row.gid)}
                              color='disabled'/>)}

                  <RemoveIcon className={classes.fnPointer}
                              onClick={this.delete.bind(this, row.gid)}
                              color='secondary'/>

                  {row.csize > 0 && (<LocalIcon className={classes.fnPointer}
                             onClick={this.local.bind(this, row.saveDir,
                               row.savePath)}
                             color='action'/>)}
                </div>
              </div>
              <div className={classes.listMeta}>
                <div style={{width: 50 + 'px'}}>{row.hSize}</div>
                <div className={classes.fnFlex1}>
                  已下载 {parseInt(row.progress, 10)}% &nbsp; 剩余 {row.eta}
                </div>
                <div>
                  {row.pieces}p/{row.conns}c/
                  {
                    row.state === 2 ? (
                      <span className={classes.ftError}>未知错误</span>) : row.speed
                  }
                </div>
              </div>
              <LinearProgress className={classes.listProgress}
                              variant="determinate"
                              value={parseInt(row.progress, 10)}/>
            </div>)
          })
        }
      </div>
    )
  }
}

Downloading.propTypes = {
  classes: PropTypes.object.isRequired,
}