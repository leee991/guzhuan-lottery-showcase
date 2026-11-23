/**
 * 接口定义 — lottery/tests/admin-security-integration.test.cjs
 * 仅保留模块导出与接口签名。
 */


const test=require('node:test'),assert=require('node:assert/strict'),fs=require('node:fs'),path=require('node:path'),os=require('node:os');const createOps=require('../lib/admin-ops'),{totp}=require('../lib/admin-security');

// Discovered route endpoints in this module (read-only contract summary):
//   - /api/admin/ops/
//   - /api/admin/ops/me
