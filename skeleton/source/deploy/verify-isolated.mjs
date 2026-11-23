/**
 * 接口定义 — source/deploy/verify-isolated.mjs
 * 仅保留模块导出与接口签名。
 */

import http from 'node:http';
import fs from 'node:fs/promises';
import os from 'node:os';
import path from 'node:path';
import {spawn} from 'node:child_process';
import assert from 'node:assert/strict';
import crypto from 'node:crypto';

async function api(route, method='GET', body, authenticated=true) {
  // 内部业务实现略
  throw new Error("not implemented");
}
