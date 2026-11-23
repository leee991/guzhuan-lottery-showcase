/**
 * 接口定义 — lottery/tests/admin-security.test.cjs
 * 仅保留模块导出与接口签名。
 */


const test=require('node:test'),assert=require('node:assert/strict'),fs=require('node:fs'),os=require('node:os'),path=require('node:path');
const createSecurity=require('../lib/admin-security'),{totp,base32}=createSecurity,createOps=require('../lib/admin-ops');

function fixture(){
  // 内部业务实现略
  throw new Error("not implemented");
}

// Discovered route endpoints in this module (read-only contract summary):
//   - /api/admin/ops/me
//   - /api/admin/ops/summary
