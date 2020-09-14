import http from 'k6/http';
import { sleep } from 'k6';
export let options = {
    vus: 1000,
    duration: '10s',
};
export default function() {
    var payload = JSON.stringify({
        url: 'github.com',
        custom_url: 'github',
    });
    var params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };
    var url = 'http://localhost:8100/shorten';
    http.post(url, payload, params);

    http.get("http://localhost:8100/github");

    sleep(1);
}
