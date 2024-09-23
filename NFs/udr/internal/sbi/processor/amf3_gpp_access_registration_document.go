/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package processor

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/udr/internal/logger"
	"github.com/free5gc/udr/internal/util"
	"github.com/free5gc/util/mongoapi"
)

func (p *Processor) AmfContext3gppProcedure(
	c *gin.Context, collName string, ueId string, patchItem []models.PatchItem,
) {
	var origValue, newValue map[string]interface{}
	var err error
	filter := bson.M{"ueId": ueId}
	if origValue, newValue, err = p.PatchDataToDBAndNotify(collName, ueId, patchItem, filter); err != nil {
		logger.DataRepoLog.Errorf("AmfContext3gppProcedure err: %+v", err)
		problemDetails := util.ProblemDetailsModifyNotAllowed("")
		c.JSON(int(problemDetails.Status), problemDetails)
	}

	PreHandleOnDataChangeNotify(ueId, CurrentResourceUri, patchItem, origValue, newValue)
	c.Status(http.StatusNoContent)
}

func (p *Processor) CreateAmfContext3gppProcedure(c *gin.Context, collName string, ueId string,
	Amf3GppAccessRegistration models.Amf3GppAccessRegistration,
) {
	filter := bson.M{"ueId": ueId}
	putData := util.ToBsonM(Amf3GppAccessRegistration)
	putData["ueId"] = ueId

	if _, err := mongoapi.RestfulAPIPutOne(collName, filter, putData); err != nil {
		logger.DataRepoLog.Errorf("CreateAmfContext3gppProcedure err: %+v", err)
	}
	c.Status(http.StatusNoContent)
}

func (p *Processor) QueryAmfContext3gppProcedure(c *gin.Context, collName string, ueId string) {
	filter := bson.M{"ueId": ueId}
	data, pd := p.GetDataFromDB(collName, filter)
	if pd != nil {
		logger.DataRepoLog.Errorf("QueryAmfContext3gppProcedure err: %s", pd.Detail)
		c.JSON(int(pd.Status), pd)
	}
	c.JSON(http.StatusOK, data)
}
