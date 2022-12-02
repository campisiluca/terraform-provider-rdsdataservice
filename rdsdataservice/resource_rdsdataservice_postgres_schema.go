package rdsdataservice

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAwsRdsdataservicePostgresSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsRdsdataservicePostgresSchemaCreate,
		Delete: resourceAwsRdsdataservicePostgresSchemaDelete,
		Exists: resourceAwsRdsdataservicePostgresSchemaExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Schema name.",
			},
			"resource_arn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DB ARN.",
			},
			"secret_arn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DBA Secret ARN.",
			},
			"database": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Schema name.",
			},
		},
	}
}

func resourceAwsRdsdataservicePostgresSchemaCreate(d *schema.ResourceData, meta interface{}) error {
	rdsdataserviceconn := meta.(*AWSClient).rdsdataserviceconn

	sql := fmt.Sprintf("CREATE SCHEMA %s;",
		d.Get("name").(string))

	createOpts := rdsdataservice.ExecuteStatementInput{
		ResourceArn: aws.String(d.Get("resource_arn").(string)),
		SecretArn:   aws.String(d.Get("secret_arn").(string)),
		Sql:         aws.String(sql),
		Database:	 aws.String(d.Get("database").(string)),
	}

	log.Printf("[DEBUG] Create Postgres Schema: %#v", createOpts)

	_, err := rdsdataserviceconn.ExecuteStatement(&createOpts)

	if err != nil {
		return fmt.Errorf("Error creating Postgres Schema: %#v", err)
	}

	d.SetId(d.Get("name").(string))
	log.Printf("[INFO] Postgres Schema ID: %s", d.Id())

	return err
}

func resourceAwsRdsdataservicePostgresSchemaDelete(d *schema.ResourceData, meta interface{}) error {
	rdsdataserviceconn := meta.(*AWSClient).rdsdataserviceconn

	sql := fmt.Sprintf("DROP SCHEMA %s;",
		d.Get("name").(string))

	createOpts := rdsdataservice.ExecuteStatementInput{
		ResourceArn: aws.String(d.Get("resource_arn").(string)),
		SecretArn:   aws.String(d.Get("secret_arn").(string)),
		Sql:         aws.String(sql),
		Database:	 aws.String(d.Get("database").(string)),
	}

	log.Printf("[DEBUG] Drop Postgres SCHEMA: %#v", createOpts)

	_, err := rdsdataserviceconn.ExecuteStatement(&createOpts)

	if err != nil {
		return fmt.Errorf("Error dropping Postgres SCHEMA: %#v", err)
	}

	d.SetId("")
	return err
}

func resourceAwsRdsdataservicePostgresSchemaExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	rdsdataserviceconn := meta.(*AWSClient).rdsdataserviceconn

	sql := fmt.Sprintf("SELECT nspname FROM pg_catalog.pg_namespace WHERE nspname='%s';",
		d.Get("name").(string))

	createOpts := rdsdataservice.ExecuteStatementInput{
		ResourceArn: aws.String(d.Get("resource_arn").(string)),
		SecretArn:   aws.String(d.Get("secret_arn").(string)),
		Sql:         aws.String(sql),
	}

	log.Printf("[DEBUG] Check Postgres Schema exists: %#v", createOpts)

	output, err := rdsdataserviceconn.ExecuteStatement(&createOpts)

	if err != nil {
		return false, fmt.Errorf("Error checking Postgres Schema exists: %#v", err)
	}

	if len(output.Records) != 1 {
		return false, nil
	}

	return true, nil
}