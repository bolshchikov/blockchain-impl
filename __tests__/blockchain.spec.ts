import { Blockchain } from "../src/blockchain";

describe("Blockchain", () => {
  it("should create a simple chain", () => {
    const bc = new Blockchain();
    bc.addBlock({ test: "test" });
    expect(bc.chain.length).toEqual(2);
    expect(bc.isChainValid()).toBe(true);
  });

  it("should invalidate the chain", () => {
    const bc = new Blockchain();
    bc.addBlock({ test: "test" });
    bc.addBlock({ test: "test2" });
    bc.chain[1].index = 80;
    expect(bc.isChainValid()).toBe(false);
  });
});
