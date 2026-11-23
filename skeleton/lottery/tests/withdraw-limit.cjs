/**
 * 接口定义 — lottery/tests/withdraw-limit.cjs
 * 仅保留模块导出与接口签名。
 */


const assert=require('node:assert/strict');
const fs=require('node:fs');
const {readSnapshot}=require('../lib/sqlite-store');

async function api(route,body){
  // 内部业务实现略
  throw new Error("not implemented");
}

// Discovered route endpoints in this module (read-only contract summary):
//   - /api/bank
//   - /api/login
//   - /api/me
//   - /api/withdraw
