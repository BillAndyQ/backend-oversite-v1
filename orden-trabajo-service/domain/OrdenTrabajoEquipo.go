package domain

import "orden-trabajo-service/utils"

type OrdenTrabajoEquipo struct {
	ID                int    `json:"id" gorm:"primaryKey"`
	NOrdenTrabajo     string `json:"n_orden_trabajo"`
	NOrdenService     string `json:"n_orden_servicio"`
	EmpresaMatriz     string `json:"empresa_matriz"`
	EmpresaSocia      string `json:"empresa_socia"`
	Ruc               string `json:"ruc"`
	FechaServicio     string `json:"fecha_servicio"`
	Certificadora     string `json:"certificadora"`
	TipoUnidad        string `json:"tipo_unidad"`
	Placa             string `json:"placa"`
	Area              string `json:"area"`
	NFactura          string `json:"n_factura"`
	FechaFactura      string `json:"fecha_factura"`
	MontoSinIGV       string `json:"monto_sin_igv"`
	MontoConIGV       string `json:"monto_con_igv"`
	IGV               string `json:"igv"`
	MontoDetraccion   string `json:"detraccion"`
	EnSoles           string `json:"en_soles"`
	EnDolares         string `json:"en_dolares"`
	Moneda            string `json:"moneda"`
	Estado            string `json:"estado"`
	FechaCreacion     string `json:"fecha_creacion"`
	FechaModificacion string `json:"fecha_modificacion"`
}

type OTEqAdminDTO struct {
	ID            int                      `json:"id" gorm:"primaryKey"`
	NOrdenTrabajo string                   `json:"n_orden_trabajo"`
	NOrdenService string                   `json:"n_orden_servicio"`
	EmpresaMatriz string                   `json:"empresa_matriz"`
	EmpresaSocia  string                   `json:"empresa_socia"`
	Ruc           string                   `json:"ruc"`
	FechaServicio utils.FechaPersonalizada `json:"fecha_servicio"`
	Certificadora string                   `json:"certificadora"`
	TipoUnidad    string                   `json:"tipo_unidad"`
	Placa         string                   `json:"placa"`
	Area          string                   `json:"area"`
}
