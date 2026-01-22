package services

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type SecretsService struct {
	client *secretsmanager.Client
}

type Secret struct {
	Name        string `json:"name" example:"prd/database"`
	Value       string `json:"value,omitempty" example:"USERNAME=admin\nPASSWORD=secret123"`
	Description string `json:"description,omitempty" example:"Production database credentials"`
}

func NewSecretsService() (*SecretsService, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := secretsmanager.NewFromConfig(cfg)
	return &SecretsService{client: client}, nil
}

func (s *SecretsService) ListSecrets(ctx context.Context) ([]Secret, error) {
	var secrets []Secret
	paginator := secretsmanager.NewListSecretsPaginator(s.client, &secretsmanager.ListSecretsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list secrets: %w", err)
		}

		for _, secret := range output.SecretList {
			secrets = append(secrets, Secret{
				Name:        aws.ToString(secret.Name),
				Description: aws.ToString(secret.Description),
			})
		}
	}

	return secrets, nil
}

func (s *SecretsService) GetSecret(ctx context.Context, name string) (*Secret, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	}

	result, err := s.client.GetSecretValue(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get secret: %w", err)
	}

	return &Secret{
		Name:  name,
		Value: aws.ToString(result.SecretString),
	}, nil
}

func (s *SecretsService) CreateSecret(ctx context.Context, secret Secret) error {
	input := &secretsmanager.CreateSecretInput{
		Name:         aws.String(secret.Name),
		SecretString: aws.String(secret.Value),
	}

	if secret.Description != "" {
		input.Description = aws.String(secret.Description)
	}

	_, err := s.client.CreateSecret(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create secret: %w", err)
	}

	return nil
}

func (s *SecretsService) UpdateSecret(ctx context.Context, name string, value string) error {
	input := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(name),
		SecretString: aws.String(value),
	}

	_, err := s.client.PutSecretValue(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to update secret: %w", err)
	}

	return nil
}
