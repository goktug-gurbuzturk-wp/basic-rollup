.PHONY: all run-anvil-default run-anvil run-tx-builder deploy-contract run-sequencer run-observer stop-anvil stop-anvil-default stop-observer

all: run-anvil-default run-anvil run-tx-builder deploy-contract run-sequencer run-observer

run-anvil-default:
	@echo "Starting Anvil node without specifying genesis or port..."
	anvil &

run-anvil:
	@echo "Starting Anvil node on port 8546 with genesis file..."
	anvil --init app/transaction-builder/example-genesis.json --port 8546 &

run-tx-builder:
	@echo "Running transaction builder..."
	cd app/transaction-builder && go run main.go &

stop-anvil:
	@echo "Stopping Anvil node on port 8546..."
	pkill -f "anvil --init app/transaction-builder/example-genesis.json --port 8546" &

deploy-contract:
	@echo "Deploying contract..."
	cd foundry && forge script script/RollupDataLayer.s.sol --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast &

run-sequencer:
	@echo "Running sequencer..."
	cd app/sequencer && go run main.go &

run-observer:
	@echo "Running observer..."
	cd app/observer && go run main.go &

stop-observer:
	@echo "Stopping observer program..."
	pkill -f "cd app/observer && go run main.go"

stop-anvil-default:
	@echo "Stopping default Anvil node..."
	pkill -f "anvil"
