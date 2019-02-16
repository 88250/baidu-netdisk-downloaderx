import React from 'react'
import PropTypes from 'prop-types'

export default class Donate extends React.Component {

  openGithub () {
    window.shell.openExternal(
      'https://github.com/b3log/baidu-netdisk-downloaderx')
  }

  render () {
    const {classes} = this.props
    return (
      <div className={classes.donate}>
        <br/><br/>
        如果你觉得 <span className={classes.link}
                    onClick={this.openGithub}>BND2</span> 还不错，欢迎成为我们的赞助者 <br/>
        <img alt="微信捐赠" className={classes.donateImg} src="wechat-donate.jpg"/>
        <img alt="支付宝捐赠" className={classes.donateImg} src="alipay-donate.jpg"/>
      </div>
    )
  }
}

Donate.propTypes = {
  classes: PropTypes.object.isRequired,
}