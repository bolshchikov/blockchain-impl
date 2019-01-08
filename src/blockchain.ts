import { Block } from "./block";

export class Blockchain {
  public chain: Block[];
  constructor() {
    this.chain = [this.createGenesisBlock()];
  }

  private createGenesisBlock(): Block {
    return new Block(0);
  }

  private getLastBlock() {
    return this.chain[this.chain.length - 1];
  }

  addBlock(data: any): void {
    const prevBlock = this.getLastBlock();
    const block = new Block(prevBlock.index + 1, data, prevBlock.hash);
    this.chain.push(block);
  }

  isChainValid(): boolean {
    for (let i = 1; i < this.chain.length; i++) {
      const currentBlock = this.chain[i];
      const previousBlock = this.chain[i - 1];

      if (currentBlock.hash !== currentBlock.calculateHash()) {
        return false;
      }

      if (currentBlock.prevHash !== previousBlock.hash) {
        return false;
      }
    }
    return true;
  }
}
