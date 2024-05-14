// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import { RollupDataLayer} from "../src/RollupDataLayer.sol";

contract RollupDataLayerScript is Script {
    function run() external {
        vm.startBroadcast();

        address sequencerAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266; // public key for first account in anvil
        console.log('sequencer address: ', sequencerAddress);
        RollupDataLayer deployedContract = new RollupDataLayer(sequencerAddress);
        console.log("Deployed RollupDataLayer at:", address(deployedContract));

        vm.stopBroadcast();
    }
}
