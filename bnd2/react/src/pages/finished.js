import React from 'react'
import PropTypes from 'prop-types'
import Button from '@material-ui/core/Button'
import RemoveIcon from '@material-ui/icons/DeleteOutline'
import LocalIcon from '@material-ui/icons/FilterCenterFocus'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import { clearCookie } from '../utils/clearCookie'
import AppBar from '@material-ui/core/AppBar'

export default class Finished extends React.Component {
  state = {
    list: [],
  }

  componentDidMount () {
    window.rws.send(JSON.stringify({cmd: 'lsctasks', param: {}}))
  }

  componentWillReceiveProps (nextProps, nextContent) {
    const data = nextProps.rwsData
    if (data.cmd === 'lsctasks') {
      this.setState({
        list: data.data,
      })
    }
  }

  componentWillUnmount () {
    window.rws.send(JSON.stringify({cmd: 'stoplsctasks', param: {}}))
  }

  deleteAll () {
    window.rws.send(JSON.stringify({cmd: 'delctaskall', param: {}}))
  }

  delete (gid) {
    window.rws.send(JSON.stringify({cmd: 'delctask', param: {gid}}))
  }

  local (saveDir, name) {
    window.shell.showItemInFolder(window.path.join(saveDir, name))
  }

  render () {
    const {classes} = this.props

    return (
      <div>
        <AppBar
          className={classes.menu}>
          <Toolbar>
            <Typography className={classes.fnFlex1} color="inherit" noWrap>
              下载完成
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
        <div className={classes.fnFlex}>
          <div className={classes.fnFlex1}>
            <Button className={classes.ftOriginal} onClick={this.deleteAll}
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
                <div className={classes.listTitle}>
                  {row.name} &nbsp;
                  <div className={classes.listMeta}> {row.hSize}</div>
                </div>
                <div>
                  <RemoveIcon className={classes.fnPointer}
                              onClick={this.delete.bind(this, row.gid)}
                              color='secondary'/>

                  <LocalIcon className={classes.fnPointer}
                             onClick={this.local.bind(this, row.saveDir,
                               row.savePath)}
                             color='action'/>
                </div>
              </div>
            </div>)
          })
        }
      </div>
    )
  }
}

Finished.propTypes = {
  classes: PropTypes.object.isRequired,
}