import  axios from 'axios'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
// // set the global baseurl
// axios.defaults.baseURL=""
// //allow to carry the cookie
// axios.defaults.withCredentials=true
const axiosApi=axios.create({
    // set the global baseurl
    baseURL:"http://localhost:8080/",
    withCredentials:true
})

axios.interceptors.request.use(function(config){
    NProgress.start();
    return config;
},error => {
    return Promise.reject(error);
})


axios.interceptors.response.use(function(config){
    NProgress.done();
    return config;
},error=>{
    NProgress.done();
    return Promise.reject(error);
})


export default  axiosApi;

