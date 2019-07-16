import React from 'react'
import PropTypes from 'prop-types'
import Button from '@material-ui/core/Button'

export default class Donate extends React.Component {

  openGithub (path) {
    window.shell.openExternal(path)
  }

  render () {
    const {classes} = this.props
    return (
      <div className={classes.donate}>
        <br/><br/>
        <h2>❤️ 欢迎成为我们的赞助者</h2>
        <br/><br/><br/>
        <span className={classes.link}
              onClick={this.openGithub.bind(this,
                'https://b3log.org/')}>B3log 开源组织</span>
        旗下包含&nbsp;
        <span className={classes.link}
              onClick={this.openGithub.bind(this,
                'https://github.com/b3log/baidu-netdisk-downloaderx')}>BND</span>、
        <span className={classes.link}
              onClick={this.openGithub.bind(this,
                'https://sym.b3log.org/')}>Symphony</span>、
        <span className={classes.link}
              onClick={this.openGithub.bind(this,
                'https://solo.b3log.org/')}>Solo</span>、
        <span className={classes.link}
              onClick={this.openGithub.bind(this,
                'https://github.com/b3log/pipe')}>Pipe</span>、
        <span className={classes.link}
              onClick={this.openGithub.bind(this,
                'https://github.com/b3log/wide')}>Wide</span>、
        <span className={classes.link}
              onClick={this.openGithub.bind(this,
                'https://github.com/b3log/latke')}>Latke</span>、
        <span className={classes.link}
              onClick={this.openGithub.bind(this,
                'https://github.com/b3log/vditor')}>Vditor</span>、
        <span className={classes.link}
              onClick={this.openGithub.bind(this,
                'https://github.com/b3log/gulu')}>Gulu</span>&nbsp;等一系列开源项目。随着项目规模的增长，我们需要有相应的资金支持才能持续项目的维护和开发。
                <br/><br/>
                如果你觉得 BND2 还算好用，可通过支付宝对我们进行赞助支持，非常感谢！
        <br/><br/><br/><br/>
        <Button className={classes.ftOriginal}
                color='primary'
                onClick={this.openGithub.bind(this,
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
        {/*<img alt="微信捐赠" className={classes.donateImg} src="wechat-donate.jpg"/>*/}
        {/*<img alt="支付宝捐赠" className={classes.donateImg} src="alipay-donate.jpg"/>*/}
      </div>
    )
  }
}

Donate.propTypes = {
  classes: PropTypes.object.isRequired,
}
