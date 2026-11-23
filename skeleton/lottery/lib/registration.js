/**
 * 接口定义 — lottery/lib/registration.js
 * 仅保留模块导出与接口签名。
 */

const crypto = require('node:crypto');
const {promisify} = require('node:util');
const renderCaptcha = require('./svg-captcha.cjs');

function allowRequest(key, maximum, windowMs) {
  // 内部业务实现略
  throw new Error("not implemented");
}

function issueCaptcha(owner, previousId) {
  // 内部业务实现略
  throw new Error("not implemented");
}

function consumeCaptcha(id, answer, owner) {
  // 内部业务实现略
  throw new Error("not implemented");
}

async function hashPassword(password) {
  // 内部业务实现略
  throw new Error("not implemented");
}

async function matchesPassword(password, stored) {
  // 内部业务实现略
  throw new Error("not implemented");
}
