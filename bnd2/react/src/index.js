import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import { ws } from './utils/ws'

(async () => {
  await ws()
  ReactDOM.render(<App/>, document.querySelector('#root'))
})()

