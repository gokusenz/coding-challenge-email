import axios from 'axios';
import { TIMEOUT_REQUEST } from '../constants/Request';

const API_URL = 'http://localhost:8080';

export default {
  getHeader() {
    const headers = {
      'Content-Type' : 'application/x-www-form-urlencoded; charset=UTF-8',
    };
    return headers;
  },

  get(path, params = {}) {
    path = API_URL + path
    return axios.get(path, {
      headers: this.getHeader(),
      params,
      timeout: TIMEOUT_REQUEST,
    });
  },

  post(path, params = {}) {
    path = API_URL + path    
    return axios.post(path,
      params,
      { headers: this.getHeader(), timeout: TIMEOUT_REQUEST },
    );
  },

  put(path, params = {}) {
    path = API_URL + path    
    return axios.put(path,
      params,
      { headers: this.getHeader(), timeout: TIMEOUT_REQUEST },
    );
  },

  delete(path, params = {}) {
    path = API_URL + path    
    return axios.delete(path, {
      headers: this.getHeader(),
      params,
      timeout: TIMEOUT_REQUEST,
    });
  },
};
