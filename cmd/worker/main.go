package main

//
//import (
//	"context"
//	"os"
//	"os/signal"
//	"time"
//
//	"github.com/red-rocket-software/reminder-go/config"
//	todoStorage "github.com/red-rocket-software/reminder-go/internal/reminder/storage"
//	"github.com/red-rocket-software/reminder-go/pkg/logging"
//	"github.com/red-rocket-software/reminder-go/pkg/postgresql"
//	"github.com/red-rocket-software/reminder-go/worker"
//)
//
//func main() {
//	cfg := config.GetConfig()
//	logger := logging.GetLogger()
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	logger.Info("Getting new db client...")
//	postgresClient, err := postgresql.NewClient(ctx, 5, *cfg)
//	if err != nil {
//		logger.Fatalf("Error create new db client:%v\n", err)
//	}
//	defer postgresClient.Close()
//
//	remindStorage := todoStorage.NewStorageTodo(postgresClient, &logger)
//
//	newWorker := worker.NewWorker(ctx, remindStorage, *cfg)
//
//	//run worker in scheduler
//	c := make(chan os.Signal, 1)
//	signal.Notify(c)
//	stop := make(chan bool)
//
//	ticker := time.NewTicker(time.Second * 5) // worker runs every 5 second
//
//	go func() {
//		defer func() { stop <- true }()
//		for {
//			select {
//			case <-ticker.C:
//				err = newWorker.ProcessSendNotification()
//				if err != nil {
//					logger.Errorf("error to process worker send notification: %v", err)
//					return
//				}
//				err = newWorker.ProcessSendDeadlineNotification()
//				if err != nil {
//					logger.Errorf("error to process worker send deadline notification: %v", err)
//					return
//				}
//			case <-stop:
//				logger.Info("closing goroutine")
//				return
//			}
//		}
//
//	}()
//	<-c
//	defer ticker.Stop()
//
//	stop <- true
//
//	<-stop
//	logger.Info("Stop application")
//}
