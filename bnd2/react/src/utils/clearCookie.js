export const clearCookie = () => {
  window.ipcRenderer.sendToHost('clearCookie')
}