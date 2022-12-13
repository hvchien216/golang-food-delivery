package subscriber

import (
	"context"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/asyncjob"
	"food_delivery/pubsub"
	"food_delivery/skio"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx   appctx.AppContext
	rtEngine skio.RealtimeEngine
}

func NewEngine(appCtx appctx.AppContext, rtEngine skio.RealtimeEngine) *consumerEngine {
	return &consumerEngine{appCtx: appCtx, rtEngine: rtEngine}
}

func (engine *consumerEngine) Start(
// rtEngine skio.RealtimeEngine,
) error {

	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		RunIncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
		EmitRealtimeAfterUserLikeRestaurant(engine.appCtx, engine.rtEngine),
	)

	engine.startSubTopic(
		common.TopicUserDislikeRestaurant,
		true,
		RunDecreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	// make color =))
	for _, item := range consumerJobs {
		log.Println("setup consumer for: ", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	// linking the pubsub system & async job system
	go func() {
		for {
			msg := <-c

			jobHldArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHldArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHldArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
