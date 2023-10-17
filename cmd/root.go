package cmd

import (
	"fmt"
	"github.com/docker/distribution/context"
	"github.com/kwanok/minsim-consumer/client"
	"github.com/kwanok/minsim-consumer/config"
	"github.com/kwanok/minsim-consumer/service/minsim"
	"github.com/kwanok/minsim-consumer/service/reddit/comment"
	"github.com/kwanok/minsim-consumer/service/reddit/submission"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kafka-consumer",
	Short: "Kafka Consumer",
	Long:  `MLOps 스터디용 카프카 컨슈머 프로그램입니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		grpcClient := client.NewGrpcClient(globalConfig)
		httpClient := client.NewHttpClient(
			&http.Client{},
			fmt.Sprintf("%s:%s", globalConfig.Inference.Host, globalConfig.Inference.Port),
		)

		commentConsumer := comment.NewRedditCommentConsumer(
			globalConfig,
			minsim.NewMinsimClient(grpcClient.ClientConn),
			httpClient,
		)

		go func() {
			err := commentConsumer.Start(ctx)
			if err != nil {
				log.Fatalf("comment consumer error: %v", err)
			}
		}()

		submissionConsumer := submission.NewRedditSubmissionConsumer(
			globalConfig,
			grpcClient,
			httpClient,
		)
		go func() {
			err := submissionConsumer.Start(ctx)
			if err != nil {
				log.Fatalf("submission consumer error: %v", err)
			}
		}()

		select {
		case <-ctx.Done():
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var cfgFile string
var globalConfig *config.Config

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func initConfig() {
	if cfgFile == "" {
		fmt.Println("config file not specified")
		os.Exit(1)
	}
	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&globalConfig); err != nil {
		fmt.Println("Can't unmarshal config:", err)
		os.Exit(1)
	}

	fmt.Println("config loaded")
}
