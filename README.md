# Noteb00k


1. Generate Protobuf (and gRPC) code 

```bash
protoc --proto_path=[DEPENDENCY_FOLDER] \
        --go_out=. --go_opt=paths=[paths=source_relative | module=MODULE] \
        --go-grpc_out=. --go-grpc_opt=[paths=source_relative | module=MODULE] \
        [PROTO_FILENAME].proto

# 
```

2. Encoding 
- Fixed-size numbers
  - 32 bits (4 bytes) or 64 bits (8 bytes)
  - e.g., fixed32, fixed64, sfixed32, sfixed64, double, float
  - fixed32 and fixed64 are binary representations
  - IEEE Standard for Floating-Point Arithmetic (IEEE 754) 

- Varints
  - Variable-length Integers
  - Add 1 as the **most significant bit(MSB)** for all the groups except the first one
  - Add 0 as the MSB for the first group 
  - e.g., int32, int64, uint32, uint64, bool
  
- Length-delimited types
  - Cannot know the lenght at compile time (dynamic arrays)
  - e.g., string, bytes

```bash
# PROTO_FILENAME: wrappers.proto
cat [INPUT_FILENAME].txt | protoc --encode=[MESSAGE_NAME] [PROTO_FILENAME].proto | hexdump -C
```

- How to choose 
  - The range of numbers needed
    - fixed, signed, and variable-length integers: [-2 ^ (NUMBER_OF_BITS - 1), 2 ^ (NUMBER_OF_BUTS - 1) - 1]
    - unsigned: [0, 2 * 2 ^ (NUMBER_OF_BITS -1) - 1]
  - The need for negative numbers
  - The data distribution