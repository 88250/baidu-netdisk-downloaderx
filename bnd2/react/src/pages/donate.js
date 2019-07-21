import React from 'react'
import PropTypes from 'prop-types'
import fetchJsonp from 'fetch-jsonp'
import Button from '@material-ui/core/Button'
import Card from '@material-ui/core/Card'
import classnames from 'classnames'
import {openURL} from '../utils/openURL'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import { clearCookie } from '../utils/clearCookie'
import AppBar from '@material-ui/core/AppBar'

export default class Donate extends React.Component {

  state = {
    payments: [],
  }

  componentDidMount () {
    fetchJsonp('https://hacpai.com/apis/sponsors').
      then(response => response.json()).
      then(json => {
        this.setState({
          payments: json.data.payments,
        })
      })
  }

  render () {
    const {classes} = this.props
    return (
      <div className={classes.donate}>
        <AppBar
          className={classes.menu}>
          <Toolbar>
            <Typography className={classes.fnFlex1} color="inherit" noWrap>
              成为赞助者
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
        <h2>❤️ 欢迎成为我们的赞助者</h2>
        如果你觉得 BND2 还算好用，请在<span className={classes.link}
      onClick={openURL.bind(this,
              'https://github.com/b3log/baidu-netdisk-downloaderx')}>项目主页</span>标星点赞并<span className={classes.link}
      onClick={openURL.bind(this,
              'https://github.com/88250')}>关注我</span>了解更多开源作品。<br/>
        也欢迎你通过支付宝进行捐赠赞助，非常感谢 🙏
        <br/><br/>
        <div className={classes.ftCenter}>
          <Button className={classes.ftOriginal}
                  color='primary'
                  onClick={openURL.bind(this,
                    'https://hacpai.com/sponsor')}
                  variant="contained">
            <svg viewBox="0 0 32 32" width="100%" height="100%"
                 className={classes.svg}>
              <path
                d="M32 21.906v-15.753c0-3.396-2.757-6.152-6.155-6.152h-19.692c-3.396 0-6.152 2.756-6.152 6.152v19.694c0 3.396 2.754 6.152 6.152 6.152h19.694c3.027 0 5.545-2.189 6.058-5.066-1.632-0.707-8.703-3.76-12.388-5.519-2.804 3.397-5.74 5.434-10.166 5.434s-7.38-2.726-7.025-6.062c0.234-2.19 1.736-5.771 8.26-5.157 3.438 0.323 5.012 0.965 7.815 1.89 0.726-1.329 1.329-2.794 1.785-4.35h-12.433v-1.233h6.151v-2.212h-7.503v-1.357h7.504v-3.195c0 0 0.068-0.499 0.62-0.499h3.077v3.692h7.999v1.357h-7.999v2.212h6.526c-0.6 2.442-1.51 4.686-2.651 6.645 1.895 0.686 10.523 3.324 10.523 3.324v0 0 0zM8.859 24.736c-4.677 0-5.417-2.953-5.168-4.187 0.246-1.227 1.6-2.831 4.201-2.831 2.987 0 5.664 0.767 8.876 2.328-2.256 2.94-5.029 4.69-7.908 4.69v0 0z"></path>
            </svg>
            &nbsp;
            使用支付宝进行赞助
          </Button>
        </div>
        <br/>
        <Card className={classes.payment}>
          {
            this.state.payments.map(row => {
              return (
                <div className={classes.listItem} key={row.oId}>
                  <div className={classnames(classes.fnFlex, classes.ft12)}>
                    {
                      row.paymentUserName
                        ? (
                          <span className={classes.link}
                                onClick={openURL.bind(this,
                                  'https://hacpai.com/member/' +
                                  row.paymentUserName)}>{row.paymentUserName}</span>)
                        : (<span className={classes.ftGray}>匿名好心人</span>)
                    }
                    <span className={classes.ftGray}>：</span>
                    <span
                      className={classnames(classes.fnFlex1,
                        classes.ftGreen)}>{row.paymentAmount}RMB</span>
                    <span className={classes.ftGray}>{row.paymentTimeStr}</span>
                  </div>
                  <div dangerouslySetInnerHTML={{ __html: row.paymentMemo}}></div>
                </div>
              )
            })
          }
        </Card>
      </div>
    )
  }
}

Donate.propTypes = {
  classes: PropTypes.object.isRequired,
}
