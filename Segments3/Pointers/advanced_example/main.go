func CreateStorageAccount(subID, string, accountName, resource_group_name string, ctx context.Context)  {
	storageAccountsClient := storage.NewAccountsClient(subID)
	storageAccountsClient.Authorizer = azureAuth()

	if azureAuth() == nil {
		log.Panicln("No Azure CLi auth detected")
	}

	future, err := storageAccountsClient.Create(
		ctx,
		resource_group_name,
		accountName,
		storageAccountsCreateParameters{
			Sku: &storage.Sku{
				Name: storage.StandardLRS},
			Kind:							: storage.Storage,
			Location:						: to.StringPtr("eastus"),
			AccoutPropertiesCreateParameters: &storage.AccoutPropertiesCreateParameters{},	
		})

	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println(future)
}