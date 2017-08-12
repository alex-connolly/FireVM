# FireVM

A stack-based virtual machine for executing smart contracts, created using Axia's [vmgen](https://github.com/end-r/vmgen).

FireVM provides a strict superset of the functionality offered by Ethereum's EVM. This means that all contracts which can be executed on the Ethereum blockchain can be executed on any chain using FireVM (although a translation mechanism is necessary).

## Advantages

### Compressed Bytecode

FireVM employs several strategies designed to minimise the size of the bytecode which must be stored on-chain.

EVM parameters are padded to 32 bytes, such that a boolean is stored as 0x0000000000000001 or 0x0000000000000000. 
This is unnecessarily wasteful, as .

FireVM stores the size of each parameter as a prefix, such that the representation of the boolean above is 0x0101 - one size byte and then one value byte.

### Improved Recursive Mapping

## Specification

See SPEC.md for full details.
