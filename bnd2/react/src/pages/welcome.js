import React from 'react'
import { Link } from 'react-router-dom'
import PropTypes from 'prop-types'
import fetchJsonp from 'fetch-jsonp'
import Card from '@material-ui/core/Card'
import CardActions from '@material-ui/core/CardActions'
import Button from '@material-ui/core/Button'
import ShareIcon from '@material-ui/icons/Share'
import PlayIcon from '@material-ui/icons/PlayArrow'
import StarIcon from '@material-ui/icons/Star'
import CardContent from '@material-ui/core/CardContent'
import Typography from '@material-ui/core/Typography'
import Snackbar from '@material-ui/core/Snackbar'
import SwipeableDrawer from '@material-ui/core/SwipeableDrawer'
import LinearProgress from '@material-ui/core/LinearProgress'
import classnames from 'classnames'
import echarts from 'echarts/lib/echarts'

require('echarts/lib/chart/line')
require('echarts/lib/component/tooltip')

export default class Welcome extends React.Component {
  state = {
    b3logList: [],
    b3log: false,
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
      }
    }
  }

  openURL (path = 'https://github.com/b3log/baidu-netdisk-downloaderx') {
    window.shell.openExternal(path)
  }

  toggleB3log = () => {
    if (this.state.b3log) {
      this.setState({b3log: false})
    } else {
      this.setState({b3log: true})
      fetchJsonp('https://hacpai.com/apis/news')
      .then(response => response.json())
      .then(json => {
          this.setState({
            b3logList: json.articles,
          })
        })
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
        <div className={classes.fnFlex}>
          <Card className={classes.welcomeCard}>
            <CardContent>
              <Typography color="textSecondary">
                完成情况
              </Typography>
              <div>
                <Typography variant="h5" component="h2">
                  {this.state.data.ctasks.ctaskCount !== 0
                    ? this.state.data.ctasks.ctaskCount + ' 个文件'
                    : '- -'}
                </Typography>
                <Typography variant="h5" component="h2">
                  {this.state.data.ctasks.ctaskCount !== 0 &&
                  this.state.data.ctasks.hTotalSize}
                </Typography>
              </div>
            </CardContent>
            {
              this.state.data.ctasks.ctaskCount === 0 &&
              (
                <CardActions>
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
                </CardActions>
              )
            }
          </Card>
          <Card
            className={classnames(classes.welcomeCard, classes.welcomeCardMid)}>
            <CardContent>
              <Typography color="textSecondary">
                点个赞支持下呗
              </Typography>
              <Typography variant="h5" component="h2">
                <span className={classes.link}
                      onClick={this.openURL.bind(this, undefined)}>BND2</span>
              </Typography>
            </CardContent>
            <CardActions>
              <Button size="small" className={classes.ftOriginal}
                      onClick={this.openURL.bind((this, undefined))}>
                <StarIcon/>
                前往 github
              </Button>
            </CardActions>
          </Card>
          <Card className={classes.welcomeCard}>
            <CardContent>
              <Typography color="textSecondary">
                随便看看
              </Typography>
              <Typography variant="h5" className={classes.link}
                          component="h2" onClick={this.toggleB3log}>
                B3log 社区动态
              </Typography>
            </CardContent>
            <CardActions>
              <Button
                className={classes.ftOriginal}
                size="small"
                onClick={this.share.bind(this)}
              >
                <ShareIcon/>
                分享 BND2
                <input className={classes.copyInput} ref={this.text}
                       readOnly="{true}"
                       value={this.state.shareText}/>
              </Button>
            </CardActions>
          </Card>
        </div>
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
        <SwipeableDrawer onClose={this.toggleB3log}
                         onOpen={this.toggleB3log} anchor="right"
                         open={this.state.b3log}>
          <div className={classes.b3log}>
            {
              this.state.b3logList.map(row => {
                return (
                  <div className={classes.listItem} key={row.articleCreateTime}
                       onClick={this.openURL.bind(this, row.articlePermalink)}>
                      <div className={classnames(classes.listTitle, classes.fnPointer)}
                      >{row.articleTitle}</div>
                  </div>
                )
              })
            }
          </div>
        </SwipeableDrawer>
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