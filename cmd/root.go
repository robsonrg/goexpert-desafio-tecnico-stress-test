package cmd

import (
	"os"

	"github.com/robsonrg/goexpert-desafio-tecnico-stress-test/internal/model"
	"github.com/robsonrg/goexpert-desafio-tecnico-stress-test/internal/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goexpert-desafio-stress-test",
	Short: "Sistema para realizar testes de carga em um serviço web",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt64("requests")
		concurrency, _ := cmd.Flags().GetInt64("concurrency")

		if requests <= 0 {
			requests = 1
		}

		if concurrency <= 0 {
			concurrency = 1
		}

		parameters := model.ParameterRequestDTO{
			Url:         url,
			Requests:    int64(requests),
			Concurrency: int64(concurrency),
		}
		usecase.Execute(parameters)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL do serviço a ser testado")
	rootCmd.Flags().Int64P("requests", "r", 1, "Número total de requests")
	rootCmd.Flags().Int64P("concurrency", "c", 1, "Número de chamadas simultâneas")
	rootCmd.MarkFlagRequired("url")
}
