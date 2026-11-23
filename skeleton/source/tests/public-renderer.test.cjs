/**
 * 接口定义 — source/tests/public-renderer.test.cjs
 * 仅保留模块导出与接口签名。
 */


const assert = require('node:assert/strict');
const fs = require('node:fs');
const vm = require('node:vm');
const path = require('node:path');

function setup(device) {
  // 内部业务实现略
  throw new Error("not implemented");
}

class NodeFixture {
  constructor(tag = 'div') {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  append(...children) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  prepend(...children) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  replaceChildren(...children) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  setAttribute(name, value) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  getAttribute(name) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  removeAttribute(name) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  addEventListener(name, handler) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  remove() {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  before(node) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  replaceWith(node) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  querySelectorAll(selector) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

  querySelector(selector) {
    // 内部业务实现略
    throw new Error("not implemented");
  }

}
