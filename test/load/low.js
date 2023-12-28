import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 5,
    duration: '5m',
};

const listNameValid = ['Obama', 'LeBron', 'Jackson']

export default () => {
    const res = http.get('http://127.0.0.1/');
    check(res, {
        'is status 200': (r) => r.status === 200,
    });
    check(res, {
        'return name': (r) => listNameValid.includes(JSON.parse(r.body)['person_name']),
    });
    sleep(1);
};
