/**
 * 接口定义 — lottery/lib/admin-security.js
 * 仅保留模块导出与接口签名。
 */


const crypto = require('node:crypto');
const fs = require('node:fs');
const path = require('node:path');
const registration = require('./registration');

function base32(bytes){
  // 内部业务实现略
  throw new Error("not implemented");
}

function decode(s){
  // 内部业务实现略
  throw new Error("not implemented");
}

function totp(secret,counter,digits=6){
  // 内部业务实现略
  throw new Error("not implemented");
}

const equal = (a,b) => {
  // 内部业务实现略
  throw new Error("not implemented");
};
