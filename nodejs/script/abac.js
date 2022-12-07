//import { newEnforcer } from 'casbin';
const { newEnforcer } = require('casbin');

async function main() {
const e =  await newEnforcer('../config/model.conf', '../config/policy.csv');

const sub = 'alice'; // 想要访问资源的用户
const obj = 'data2'; // 将要被访问的资源
const act = 'read'; // 用户对资源进行的操作

const res = await e.enforce(sub, obj, act);
if (res) {
    // 允许alice读取data1
    console.log('permit alice to read data1');
} else {
    // 拒绝请求，抛出异常
    console.log('deny the request');
}
}

main();
