/**
 * 接口定义 — lottery/tests/database.test.cjs
 * 仅保留模块导出与接口签名。
 */

const test = require('node:test');
const assert = require('node:assert/strict');
const fs = require('node:fs/promises');
const path = require('node:path');
const os = require('node:os');
const createDatabase = require('../lib/database');
