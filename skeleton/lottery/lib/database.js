/**
 * 接口定义 — lottery/lib/database.js
 * 仅保留模块导出与接口签名。
 */

const fs = require('node:fs/promises');
const path = require('node:path');
const crypto = require('node:crypto');
const {AsyncLocalStorage} = require('node:async_hooks');
