// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract RollupDataLayer {
    address public sequencer;
    bytes[] public transactions;

    constructor(address _sequencer) {
        sequencer = _sequencer;
    }

    function addRollupTransaction(bytes calldata transaction) external {
        require(msg.sender == sequencer, "Only the sequencer can add transactions");
        transactions.push(transaction);
    }
}
