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

func (p *Processor) AmfContextNon3gppProcedure(
	c *gin.Context, ueId string, collName string, patchItem []models.PatchItem,
	filter bson.M,
) {
	var err error
	var origValue, newValue map[string]interface{}
	if origValue, newValue, err = p.PatchDataToDBAndNotify(collName, ueId, patchItem, filter); err != nil {
		logger.DataRepoLog.Errorf("AmfContextNon3gppProcedure err: %+v", err)
		pd := util.ProblemDetailsSystemFailure(err.Error())
		c.JSON(int(pd.Status), pd)
	}
	PreHandleOnDataChangeNotify(ueId, CurrentResourceUri, patchItem, origValue, newValue)
	c.Status(http.StatusOK)
}

func (p *Processor) CreateAmfContextNon3gppProcedure(
	c *gin.Context, AmfNon3GppAccessRegistration models.AmfNon3GppAccessRegistration,
	collName string, ueId string,
) {
	putData := util.ToBsonM(AmfNon3GppAccessRegistration)
	putData["ueId"] = ueId
	filter := bson.M{"ueId": ueId}

	if _, err := mongoapi.RestfulAPIPutOne(collName, filter, putData); err != nil {
		logger.DataRepoLog.Errorf("CreateAmfContextNon3gppProcedure err: %+v", err)
	}

	c.Data(http.StatusNoContent, "application/json", nil)
}

func (p *Processor) QueryAmfContextNon3gppProcedure(c *gin.Context, collName string, ueId string) {
	filter := bson.M{"ueId": ueId}
	data, pd := p.GetDataFromDB(collName, filter)
	if pd != nil {
		logger.DataRepoLog.Errorf("QueryAmfContextNon3gppProcedure err: %s", pd.Detail)
		c.JSON(int(pd.Status), pd)
		return
	}
	c.JSON(http.StatusOK, data)
}
