import { setMonth } from '@/api/admin'

const actions = {
  // post month
  month({ commit }, month) {
    return new Promise((resolve, reject) => {
      setMonth({ month: month }).then(response => {
        const { data } = response
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  actions
}

