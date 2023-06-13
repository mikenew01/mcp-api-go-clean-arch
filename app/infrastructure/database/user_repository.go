package database

import (
	"clean-arch/adapter/repository"
	"clean-arch/infrastructure/logger"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const tableName = "users"

type DynamoDBClient struct {
	awsRegion   string
	awsEndpoint string
	client      *dynamodb.Client
}

var _ repository.NoSQL = (*DynamoDBClient)(nil)

func NewDynamoDBClient(awsRegion string, awsEndpoint string) *DynamoDBClient {
	dbClient := DynamoDBClient{
		awsRegion:   awsRegion,
		awsEndpoint: awsEndpoint,
	}
	dbClient.client = dbClient.loadDynamoDBClient()
	return &dbClient
}

func (client DynamoDBClient) loadDynamoDBClient() *dynamodb.Client {
	logger.StructuredLog("DynamoDBClient", "loadDynamoDBClient", "", "").Info("Init load dynamodb client")
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
			if client.awsEndpoint != "" {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           client.awsEndpoint,
					SigningRegion: client.awsRegion,
				}, nil
			}
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})),
		config.WithRegion(client.awsRegion),
	)

	if err != nil {
		panic(err)
	}

	logger.StructuredLog("DynamoDBClient", "loadDynamoDBClient", "", "").Info("Finish load dynamodb client")
	return dynamodb.NewFromConfig(awsConfig, func(opt *dynamodb.Options) {
		opt.Region = awsConfig.Region
	})
}

func (dynamodbClient DynamoDBClient) FindById(ctx context.Context, id string) (repository.UserEntity, error) {
	logger.StructuredLog("DynamoDBClient", "FindById", "input", id).Info("Init get user function")
	out, err := dynamodbClient.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return repository.UserEntity{}, err
	}

	var output repository.UserEntity
	err = attributevalue.UnmarshalMap(out.Item, &output)
	if err != nil {
		return repository.UserEntity{}, err
	}

	logger.StructuredLog("DynamoDBClient", "FindById", "output", output).Info("Finish get user function")
	return output, nil
}

func (dynamodbClient DynamoDBClient) PutItem(ctx context.Context, item repository.UserEntity) (repository.UserEntity, error) {
	logger.StructuredLog("DynamoDBClient", "PutItem", "input", item).Info("Init create user function")
	_, err := dynamodbClient.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"id":    &types.AttributeValueMemberS{Value: item.Id},
			"name":  &types.AttributeValueMemberS{Value: item.Name},
			"email": &types.AttributeValueMemberS{Value: item.Email},
		},
	})

	if err != nil {
		return repository.UserEntity{}, err
	}

	logger.StructuredLog("DynamoDBClient", "PutItem", "output", item).Info("Finish create user function")
	return item, nil
}

func (dynamodbClient DynamoDBClient) UpdateItem(ctx context.Context, item repository.UserEntity) (repository.UserEntity, error) {
	logger.StructuredLog("DynamoDBClient", "UpdateItem", "input", item).Info("Init update user function")
	_, err := dynamodbClient.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: item.Id},
		},
		UpdateExpression: aws.String("set #user_name = :name, email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name":  &types.AttributeValueMemberS{Value: item.Name},
			":email": &types.AttributeValueMemberS{Value: item.Email},
		},
		ExpressionAttributeNames: map[string]string{
			"#user_name": "name",
		},
	})

	if err != nil {
		return repository.UserEntity{}, err
	}

	logger.StructuredLog("DynamoDBClient", "UpdateItem", "output", item).Info("Finish update user function")
	return item, nil
}

func (dynamodbClient DynamoDBClient) DeleteItem(ctx context.Context, id string) error {
	logger.StructuredLog("DynamoDBClient", "DeleteItem", "input", id).Info("Init delete user function")
	_, err := dynamodbClient.client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return err
	}

	logger.StructuredLog("DynamoDBClient", "DeleteItem", "output", id).Info("Finish delete user function")
	return nil
}
