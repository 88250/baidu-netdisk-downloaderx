import side from './side'
import list from './list'
import utils from './utils'
import donate from './donate'
import welcome from './welcome'

export default theme => (
  Object.assign({
    b3log: {
      border: 0,
      height: '100vh',
      width: '320px',
    },
    content: {
      fontFamily: 'Helvetica Neue,Luxi Sans,DejaVu Sans,Tahoma,Hiragino Sans GB,Microsoft Yahei,sans-serif',
      padding: '63px 24px 24px 24px',
      height: '100vh',
      position: 'absolute',
      left: '75px',
      right: 0,
      overflow: 'auto',
    },
    menu: {
      left: '75px',
      width: 'auto',
      '&>div': {
        minHeight: '49px'
      }
    },
  }, utils, side, list, donate, welcome)
)
