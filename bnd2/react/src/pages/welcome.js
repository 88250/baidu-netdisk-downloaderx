import React from 'react'
import { Link } from 'react-router-dom'
import PropTypes from 'prop-types'
import classnames from 'classnames'
import Card from '@material-ui/core/Card'
import CardActions from '@material-ui/core/CardActions'
import Button from '@material-ui/core/Button'
import ShareIcon from '@material-ui/icons/Share'
import PlayIcon from '@material-ui/icons/PlayArrow'
import GetAPPIcon from '@material-ui/icons/GetApp'
import CardContent from '@material-ui/core/CardContent'
import Typography from '@material-ui/core/Typography'
import Snackbar from '@material-ui/core/Snackbar'
import LinearProgress from '@material-ui/core/LinearProgress'
import Toolbar from '@material-ui/core/Toolbar'
import AppBar from '@material-ui/core/AppBar'
import Divider from '@material-ui/core/Divider'
import echarts from 'echarts/lib/echarts'
import { openURL } from '../utils/openURL'
import { clearCookie } from '../utils/clearCookie'
import fetchJsonp from 'fetch-jsonp'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableRow from '@material-ui/core/TableRow'
import TableCell from '@material-ui/core/TableCell'

require('echarts/lib/chart/line')
require('echarts/lib/component/tooltip')

export default class Welcome extends React.Component {
  state = {
    b3logList: [],
    showB3log: 'block',
    copied: false,
    shareText: 'https://hacpai.com/tag/bnd',
    data: {
      ctasks: {totalSize: 0, hTotalSize: '0 B', ctaskCount: 0},
      tasks: {
        totalSize: 0,
        currentSize: 0,
        hTotalSize: '0 B',
        hCurrentSize: '0 B',
        progress: 0,
        speed: '0 B',
        taskCount: 0,
      },
      speeds: [],
    },
  }

  constructor (props) {
    super(props)
    this.text = React.createRef()
    this.chart = {}
    this.chartOption = {
      animation: false,
      legend: {
        show: false,
      },
      tooltip: {
        axisPointer: {
          lineStyle: {
            color: '#eee',
          },
        },
        trigger: 'axis',
        formatter: '下载速度 {b0}',
      },
      xAxis: [
        {
          axisLine: {
            lineStyle: {
              color: '#eee',
            },
          },
          axisTick: {
            show: false,
          },
          axisLabel: {
            color: '#fff',
          },
          type: 'category',
          data: [0],
        },
      ],
      yAxis: [
        {
          splitNumber: 3,
          axisLine: {
            lineStyle: {
              color: '#fff',
            },
          },
          axisTick: {
            show: false,
          },
          splitLine: {
            lineStyle: {
              type: 'dashed',
              color: 'rgba(35, 81, 83 ,.7)',
            }
            ,
            show: true,
          },
          axisLabel: {
            color: '#fff',
          },
        },
      ],
      series: {
        name: '下载速度',
        type: 'line',
        symbolSize: 5,
        smooth: true,
        sampling: 'average',
        itemStyle: {
          normal: {
            color: '#679df6',
          },
        },
        areaStyle: {
          normal: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: '#679df6',
              }, {
                offset: 1,
                color: 'rgba(255, 255, 255, 0.3)',
              }]),
          },
        },
        data: [0],
      },
    }
  }

  componentDidMount () {
    this.chart = echarts.init(document.getElementById('chart'))
    document.getElementById('chart').style.width = (window.outerWidth - 124) +
      'px'
    window.rws.send(JSON.stringify({cmd: 'statistic', param: {}}))

    fetchJsonp('https://hacpai.com/apis/news').
      then(response => response.json()).
      then(json => {
        this.setState({
          b3logList: json.articles,
        })
      })
  }

  componentWillUnmount () {
    window.rws.send(JSON.stringify({cmd: 'stopstatistic', param: {}}))
  }

  componentWillReceiveProps (nextProps, nextContent) {
    const data = nextProps.rwsData
    if (data.cmd === 'statistic') {
      this.setState({
        data: data.data,
      })
      if (data.data.tasks.taskCount !== 0) {
        this.chartOption.series.data = data.data.speeds || []
        this.chartOption.xAxis[0].data = data.data.hSpeeds || []
        this.chart.setOption(this.chartOption)
        this.setState({showB3log: 'none'})
      } else {
        this.setState({showB3log: 'block'})
      }
    }
  }

  share = () => {
    this.text.current.select()
    document.execCommand('copy')
    this.setState({
      copied: true,
    })

    setTimeout(() => {
      this.setState({
        copied: false,
      })
    }, 2000)
  }

  render () {
    const {classes} = this.props
    return (
      <div>
        <AppBar
          className={classes.menu}>
          <Toolbar>
            <Typography className={classes.fnFlex1} color="inherit" noWrap>
              欢迎使用
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
          <Card className={classes.welcomeCard}>
            <CardContent>
              <Typography color="textSecondary">
                完成情况
              </Typography>
              <Typography variant="h5" component="h2">
                <Link to='/finished'
                      className={classes.link}>{this.state.data.ctasks.ctaskCount} 个文件</Link>
              </Typography>
            </CardContent>
            <CardActions>
              {
                this.state.data.ctasks.ctaskCount === 0 ?
                  (
                    <Link to='/index' className={classes.ftOriginal}>
                      <Button
                        className={classes.ftOriginal}
                        size="small"
                        onClick={this.share.bind(this)}
                      >
                        <PlayIcon/>
                        查看全部文件
                      </Button>
                    </Link>
                  )
                  :
                  (<Link to='/finished' className={classes.ftOriginal}>
                      <Button size="small" className={classes.ftOriginal}>
                        <GetAPPIcon/>
                        {this.state.data.ctasks.hTotalSize}
                      </Button>
                    </Link>
                  )
              }
            </CardActions>
          </Card>
          <Card
            className={classnames(classes.welcomeCard, classes.welcomeCardMid)}>
            <CardContent>
              <Typography color="textSecondary">
                点个赞支持下呗
              </Typography>
              <Typography variant="h5" component="h2">
                <span className={classes.link}
                      onClick={openURL.bind(this,
                        'https://github.com/b3log/baidu-netdisk-downloaderx')}>前往 GitHub</span>
              </Typography>
            </CardContent>
            <CardActions>
              <Button
                className={classes.ftOriginal}
                size="small"
                onClick={this.share.bind(this)}
              >
                <ShareIcon/>
                分享 BND
                <input className={classes.copyInput} ref={this.text}
                       readOnly="{true}"
                       value={this.state.shareText}/>
              </Button>
            </CardActions>
          </Card>
          <Card className={classes.welcomeCard}>
            <CardContent>
              <Typography color="textSecondary">
                随便逛逛
              </Typography>
              <Typography variant="h5" className={classes.link}
                          component="h2" onClick={openURL.bind(this,
                'https://github.com/b3log/30-seconds-zh_CN')}>
                前端知识精选集
              </Typography>
            </CardContent>
            <CardActions>
              <Button size="small" className={classes.ftOriginal}
                      onClick={openURL.bind(this,
                        'https://hacpai.com/domain/frontend')}>
                📙 浏览前端相关讨论
              </Button>
            </CardActions>
          </Card>
        </div>

        <Card style={{display: this.state.showB3log}}>
          <Table className={classes.ftOriginal}>
            <TableBody>
              {
                this.state.b3logList.map(row => {
                  return (
                    <TableRow
                      hover
                      key={row.articleCreateTime}
                    >
                      <TableCell className={classes.link}
                                 onClick={openURL.bind(this,
                                   row.articlePermalink)}>
                        {row.articleTitle}
                      </TableCell>
                    </TableRow>
                  )
                })
              }
            </TableBody>
          </Table>
        </Card>

        {
          this.state.data.tasks.taskCount !== 0 &&
          (
            <Card>
              <CardContent>
                <br/>
                <div className={classes.listMeta}>
                  <div style={{
                    width: 50 + 'px',
                  }}>{this.state.data.tasks.hTotalSize}</div>
                  <div className={classes.fnFlex1}>
                    已下载 {parseInt(this.state.data.tasks.progress, 10)}%
                  </div>
                  <div>{this.state.data.tasks.hSpeed}/s</div>
                </div>
                <LinearProgress className={classes.welcomeProcess}
                                variant="determinate"
                                value={parseInt(this.state.data.tasks.progress,
                                  10)}/>
                <br/>
              </CardContent>
            </Card>
          )
        }
        <Card className={this.state.data.tasks.taskCount === 0
          ? classes.welcomeChartHide
          : classes.welcomeChart}>
          <div id="chart" className={classes.welcomeChartContent}></div>
        </Card>

        <Snackbar
          anchorOrigin={{vertical: 'top', horizontal: 'center'}}
          open={this.state.copied}
          message="已复制"
        />
      </div>
    )
  }
}

Welcome.propTypes = {
  classes: PropTypes.object.isRequired,
}
