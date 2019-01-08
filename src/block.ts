import { SHA256 } from "crypto-js";

export class Block {
  private timestamp: number;
  private nonce: number = 0;
  public hash: string;

  constructor(public index = 0, private data = {}, public prevHash = "") {
    this.timestamp = Date.now();
    this.hash = this.calculateHash();
  }

  mine(difficulty: number) {
    while (
      this.hash.substring(0, difficulty) !== Array(difficulty + 1).join("0")
    ) {
      this.nonce++;
      this.hash = this.calculateHash();
    }
    console.log("BLOCK MINED: " + this.hash);
  }

  calculateHash() {
    return SHA256(
      `${this.index}${this.timestamp}${this.prevHash}${JSON.stringify(
        this.data
      )}${this.nonce}`
    ).toString();
  }
}
