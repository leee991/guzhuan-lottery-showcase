/**
 * 接口定义 — lottery/tests/sqlite.test.cjs
 * 仅保留模块导出与接口签名。
 */

const test = require('node:test');
const assert = require('node:assert/strict');
const fs = require('node:fs');
const os = require('node:os');
const path = require('node:path');
const createStore = require('../lib/sqlite-store');
const createDatabase = require('../lib/database');
