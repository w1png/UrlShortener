package graphqlHandlers

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/utils"

	pb "github.com/w1png/urlshortener/pkg/url/proto"
	"google.golang.org/grpc/status"
)

func getSchema() (graphql.Schema, error) {
  urlType := models.Url{}.GraphQLType()

  query := graphql.Fields{
    "url": &graphql.Field{
      Type: urlType,
      Args: graphql.FieldConfigArgument{
        "alias": &graphql.ArgumentConfig{
          Type: graphql.String,
        },
      },
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        alias, ok := p.Args["alias"].(string)
        if !ok {
          return nil, nil
        }

        client := pb.NewUrlServiceClient(utils.UrlGRPCConnection)
        resp, err := client.GetUrl(context.Background(), &pb.GetRequest{Alias: alias})
        if err != nil {
          return nil, status.Error(404, err.Error())
        }

        return resp, nil
      },
    },
  }

  mutation := graphql.Fields{
    "createUrl": &graphql.Field{
      Type: urlType,
      Args: graphql.FieldConfigArgument{
        "url": &graphql.ArgumentConfig{
          Type: graphql.String,
        },
      },
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        url, ok := p.Args["url"].(string)
        if !ok {
          return nil, nil
        }
        
        client := pb.NewUrlServiceClient(utils.UrlGRPCConnection)
        resp, err := client.CreateUrl(context.Background(), &pb.CreateRequest{Url: url})
        if err != nil {
          return nil, status.Error(500, err.Error())
        }
      
        return resp, nil
      },
    },
  }

  schemaConfig := graphql.SchemaConfig{
    Query: graphql.NewObject(graphql.ObjectConfig{
      Name: "Query",
      Fields: query,
    }),
    Mutation: graphql.NewObject(graphql.ObjectConfig{
      Name: "Mutation",
      Fields: mutation,
    }),
  }

  return graphql.NewSchema(schemaConfig)
}

func GetHTTPHandler() (handler.Handler, error) {
  schema, err := getSchema()
  if err != nil {
    return handler.Handler{}, err
  }

  return *handler.New(&handler.Config{
    Schema: &schema,
    Pretty: true,
    GraphiQL: true,
  }), nil
}
