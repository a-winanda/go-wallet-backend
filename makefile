mock-interface:
	/home/alif.winanda/Documents/go/bin/mockery --dir=./service/ --name=user_service --output=./mocks

mock-user-repo:
	/home/alif.winanda/Documents/go/bin/mockery --dir=./repository/ --name=UserRepository --output=./mocks


mock-transaction-repo:
	/home/alif.winanda/Documents/go/bin/mockery --dir=./repository/ --name=TransactionRepository --output=./mocks