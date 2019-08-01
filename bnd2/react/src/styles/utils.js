import DEFAULT_FONT from './font'

export default {
  fnHide: {
    display: 'none',
  },
  fnFlex: {
    display: 'flex',
  },
  fnFlex1: {
    flex: 1,
  },
  ft12: {
    fontSize: '12px',
  },
  ftNormalSize: {
    fontSize: '0.875rem'
  },
  ftYellow: {
    color: '#ffcd38',
  },
  fnPointer: {
    cursor: 'pointer',
  },
  ftGray: {
    color: 'rgba(0,0,0,.54)',
  },
  ftGreen: {
    color: '#569e3d',
  },
  ftPurple: {
    color: '#563d7c',
  },
  ftError: {
    color: '#d23f31',
  },
  ftCenter: {
    textAlign: 'center',
  },
  ftBreak: {
    wordBreak: 'break-all',
  },
  link: {
    color: '#4285f4',
    textDecoration: 'none',
    cursor: 'pointer',
    wordBreak: 'break-all',
    '&:hover': {
      textDecoration: 'underline',
    },
  },
  copyInput: {
    color: '#4285f4',
    height: 0,
    width: '1px',
    margin: 0,
    border: 0,
    padding: 0,
    position: 'absolute',
  },
  dialog: {
    width: '360px',
    wordBreak: 'break-all',
  },
  ftOriginal: {
    fontFamily: DEFAULT_FONT,
    textDecoration: 'none',
    textTransform: 'inherit',
    '&:hover': {
      textDecoration: 'none',
    },
  },
}