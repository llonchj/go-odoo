package cmd

import (
	"fmt"

	"../tmpl"

	"github.com/spf13/cobra"
)

func upsert(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("At least one argument needed")
		return
	}
	c, err := getClient(cmd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if args[0] == "all" {
		args, err = c.GetAllModels()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	for _, arg := range args {
		var fakeContent map[string]map[string]interface{}
		err = c.DoRequest("fields_get", arg, []interface{}{}, map[string][]string{"attributes": []string{"type"}}, &fakeContent)
		if err != nil {
			if err.Error() != "error: \"\" code: 2" {
				fmt.Println(err.Error())
				return
			}
		}
		if len(fakeContent) == 0 {
			fmt.Println(fmt.Sprintf("WARN: The model %s was not found", arg))
			continue
		}
		err = tmpl.UpsertTypes(resolveContent(fakeContent), arg)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = tmpl.UpsertAPI(arg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func resolveContent(fakeContent map[string]map[string]interface{}) map[string]map[string]string {
	content := make(map[string]map[string]string, len(fakeContent))
	for modelName, field := range fakeContent {
		for fieldType, fieldContent := range field {
			if fieldType == "type" {
				content[modelName] = map[string]string{fieldType: fieldContent.(string)}
			}
		}
	}
	return content
}
