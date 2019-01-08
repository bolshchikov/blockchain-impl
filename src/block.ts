import { SHA256 } from "crypto-js";

export class Block {
  private timestamp: number;
  public hash: string;
  constructor(public index = 0, private data = {}, public prevHash = "") {
    this.timestamp = Date.now();
    this.hash = this.calculateHash();
  }
  calculateHash() {
    return SHA256(
      `${this.index}${this.timestamp}${this.prevHash}${JSON.stringify(
        this.data
      )}`
    ).toString();
  }
}
