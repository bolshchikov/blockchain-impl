import { Blockchain } from "./blockchain";

const chain = new Blockchain();
chain.addBlock({ amount: 4 });
chain.addBlock({ amount: 8 });

console.log(chain.isChainValid());

// try to temper
chain.chain[1].index = 80;
console.log(chain.isChainValid());
