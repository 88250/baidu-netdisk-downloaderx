import side from './side'
import list from './list'
import utils from './utils'
import donate from './donate'
import welcome from './welcome'
import DEFAULT_FONT from './font'

export default theme => (
  Object.assign({
    content: {
      fontFamily: DEFAULT_FONT,
      padding: '63px 24px 24px 24px',
      height: '100vh',
      position: 'absolute',
      left: '58px',
      right: 0,
      overflow: 'auto',
    },
    menu: {
      left: '58px',
      width: 'auto',
      '&>div': {
        minHeight: '51px',
      }
    },
  }, utils, side, list, donate, welcome)
)
