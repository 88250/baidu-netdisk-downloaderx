import ReconnectingWebSocket from 'reconnectingwebsocket'

export const ws = () => {
  return new Promise(function (resolve) {
    if (!window.rws) {
      window.rws = new ReconnectingWebSocket('ws://localhost:6804/ws', null,
        {reconnectInterval: 10000})
    }

    window.rws.onopen = function () {
      resolve()
    }
  })
}