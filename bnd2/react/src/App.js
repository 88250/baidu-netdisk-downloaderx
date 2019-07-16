import React from 'react'
import { BrowserRouter, Route, Link } from 'react-router-dom'
import PropTypes from 'prop-types'
import classnames from 'classnames'
import Drawer from '@material-ui/core/Drawer'
import Badge from '@material-ui/core/Badge'
import { withStyles } from '@material-ui/core/styles'
import List from '@material-ui/core/List'
import Divider from '@material-ui/core/Divider'
import ListItem from '@material-ui/core/ListItem'
import ListItemIcon from '@material-ui/core/ListItemIcon'
import InboxIcon from '@material-ui/icons/MoveToInbox'
import AllInboxIcon from '@material-ui/icons/AllInbox'
import DoneOutlineIcon from '@material-ui/icons/DoneOutline'
import ChartIcon from '@material-ui/icons/BubbleChart'
import FavoriteIcon from '@material-ui/icons/FavoriteBorder'
import withRoot from './withRoot'
import Index from './pages/index'
import Downloading from './pages/downloading'
import Finished from './pages/finished'
import Donate from './pages/donate'
import Welcome from './pages/welcome'
import styles from './styles'

class App extends React.Component {
  state = {
    count: 0,
    downloadingData: '',
    finishedData: '',
    indexData: '',
    welcomeData: '',
  }

  componentDidMount () {
    window.rws.send(JSON.stringify({cmd: 'counttasks', param: {}}))
    window.rws.onmessage = (evt) => {
      const data = JSON.parse(evt.data)
      if (data.cmd === 'counttasks') {
        this.setState({
          count: data.data.taskCount,
        })
      }
      if (data.cmd === 'statistic') {
        this.setState({
          welcomeData: data,
        })
      }
      if (data.cmd === 'lstasks') {
        this.setState({
          downloadingData: data,
        })
      }
      if (data.cmd === 'traverse' || data.cmd === 'ls') {
        this.setState({
          indexData: data,
        })
      }
      if (data.cmd === 'lsctasks') {
        this.setState({
          finishedData: data,
        })
      }

    }
  }

  render () {
    const {classes} = this.props
    return (
      <BrowserRouter>
        <div>
          <Drawer variant="permanent">
            <List className={classes.sideLi}>
              <Route exact
                     path="/"
                     children={({match}) => (
                       <Link className={match && classes.sideItemCurrent}
                             to="/">
                         <ListItem button>
                           <ListItemIcon>
                             <ChartIcon className={classnames(classes.sideSVG,
                               classes.ftYellow)}/>
                           </ListItemIcon>
                         </ListItem>
                       </Link>
                     )}
              />
              <Divider/>
              <Route
                path="/index"
                children={({match}) => (
                  <Link className={match && classes.sideItemCurrent}
                        to="/index">
                    <ListItem button>
                      <ListItemIcon>
                        <AllInboxIcon color='primary'
                                      className={classes.sideSVG}/>
                      </ListItemIcon>
                    </ListItem>
                  </Link>
                )}
              />
              <Divider/>
              <Route
                path="/downloading"
                children={({match}) => (
                  <Link className={match && classes.sideItemCurrent}
                        to="/downloading">
                    <ListItem button>
                      <ListItemIcon>
                        {
                          this.state.count !== 0 ? (
                            <Badge badgeContent={this.state.count}
                                   color="secondary">
                              <InboxIcon className={classnames(classes.sideSVG,
                                classes.ftPurple)}/>
                            </Badge>) : (
                            <InboxIcon className={classnames(classes.sideSVG,
                              classes.ftPurple)}/>)
                        }
                      </ListItemIcon>
                    </ListItem>
                  </Link>
                )}
              />
              <Route
                path="/finished"
                children={({match}) => (
                  <Link className={match && classes.sideItemCurrent}
                        to="/finished">
                    <ListItem button>
                      <ListItemIcon>
                        <DoneOutlineIcon className={classnames(classes.sideSVG,
                          classes.ftGreen)}/>
                      </ListItemIcon>
                    </ListItem>
                  </Link>
                )}
              />
              <Divider/>
              <Route
                path="/donate"
                children={({match}) => (
                  <Link className={match && classes.sideItemCurrent}
                        to="/donate">
                    <ListItem button>
                      <ListItemIcon>
                        <FavoriteIcon color='secondary'
                                      className={classes.sideSVG}/>
                      </ListItemIcon>
                    </ListItem>
                  </Link>
                )}
              />
            </List>
          </Drawer>
          <main className={classes.content}>
            <Route path="/" exact render={props => (
              <Welcome classes={classes} rwsData={this.state.welcomeData}/>
            )}/>
            <Route path="/index" render={props => (
              <Index classes={classes} rwsData={this.state.indexData}/>
            )}/>
            <Route path="/downloading" render={props => (
              <Downloading classes={classes}
                           rwsData={this.state.downloadingData}/>
            )}/>
            <Route path="/finished" render={props => (
              <Finished classes={classes} rwsData={this.state.finishedData}/>
            )}/>
            <Route path="/donate" render={props => (
              <Donate classes={classes}/>
            )}/>
          </main>
        </div>
      </BrowserRouter>
    )
  }
}

App.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withRoot(withStyles(styles)(App))
