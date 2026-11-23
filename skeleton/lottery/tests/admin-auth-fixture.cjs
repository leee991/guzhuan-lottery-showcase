/**
 * 接口定义 — lottery/tests/admin-auth-fixture.cjs
 * 仅保留模块导出与接口签名。
 */


const fs=require('node:fs'),os=require('node:os'),path=require('node:path');const {totp}=require('../lib/admin-security');

// Discovered route endpoints in this module (read-only contract summary):
//   - /api/admin/login
//   - /api/admin/ops/mfa/verify
