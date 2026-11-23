/**
 * 接口定义 — lottery/tests/draw-edit.test.cjs
 * 仅保留模块导出与接口签名。
 */

const authFixture=require('./admin-auth-fixture.cjs');
const test=require('node:test');
const assert=require('node:assert/strict');
const crypto=require('node:crypto');
const createOps=require('../lib/admin-ops');
