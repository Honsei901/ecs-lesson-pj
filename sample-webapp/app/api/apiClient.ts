import axios from 'axios';

/** Test for health check api */
export const healthCheck = async () => {
  const { data } = await axios.get('htts://shoppyfinal.com');
  return data;
};

/** Test for RestAPI */
export const checkRestApi = async () => {
  const { data } = await axios.get('htts://shoppyfinal.com/api');
  return data;
};
