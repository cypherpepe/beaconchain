package dataaccess

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/gobitfly/beaconchain/pkg/api/enums"
	t "github.com/gobitfly/beaconchain/pkg/api/types"
	"github.com/gobitfly/beaconchain/pkg/commons/utils"
	"github.com/lib/pq"
)

type NotificationsRepository interface {
	GetNotificationOverview(ctx context.Context, userId uint64) (*t.NotificationOverviewData, error)

	GetDashboardNotifications(ctx context.Context, userId uint64, chainIds []uint64, cursor string, colSort t.Sort[enums.NotificationDashboardsColumn], search string, limit uint64) ([]t.NotificationDashboardsTableRow, *t.Paging, error)
	// depending on how notifications are implemented, we may need to use something other than `notificationId` for identifying the notification
	GetValidatorDashboardNotificationDetails(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, epoch uint64) (*t.NotificationValidatorDashboardDetail, error)
	GetAccountDashboardNotificationDetails(ctx context.Context, dashboardId uint64, groupId uint64, epoch uint64) (*t.NotificationAccountDashboardDetail, error)

	GetMachineNotifications(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationMachinesColumn], search string, limit uint64) ([]t.NotificationMachinesTableRow, *t.Paging, error)
	GetClientNotifications(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationClientsColumn], search string, limit uint64) ([]t.NotificationClientsTableRow, *t.Paging, error)
	GetRocketPoolNotifications(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationRocketPoolColumn], search string, limit uint64) ([]t.NotificationRocketPoolTableRow, *t.Paging, error)
	GetNetworkNotifications(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationNetworksColumn], limit uint64) ([]t.NotificationNetworksTableRow, *t.Paging, error)

	GetNotificationSettings(ctx context.Context, userId uint64) (*t.NotificationSettings, error)
	UpdateNotificationSettingsGeneral(ctx context.Context, userId uint64, settings t.NotificationSettingsGeneral) error
	UpdateNotificationSettingsNetworks(ctx context.Context, userId uint64, chainId uint64, settings t.NotificationSettingsNetwork) error
	UpdateNotificationSettingsPairedDevice(ctx context.Context, userId uint64, pairedDeviceId string, name string, IsNotificationsEnabled bool) error
	DeleteNotificationSettingsPairedDevice(ctx context.Context, userId uint64, pairedDeviceId string) error
	UpdateNotificationSettingsClients(ctx context.Context, userId uint64, clientId uint64, IsSubscribed bool) (*t.NotificationSettingsClient, error)
	GetNotificationSettingsDashboards(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationSettingsDashboardColumn], search string, limit uint64) ([]t.NotificationSettingsDashboardsTableRow, *t.Paging, error)
	UpdateNotificationSettingsValidatorDashboard(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, settings t.NotificationSettingsValidatorDashboard) error
	UpdateNotificationSettingsAccountDashboard(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, settings t.NotificationSettingsAccountDashboard) error
}

func (d *DataAccessService) GetNotificationOverview(ctx context.Context, userId uint64) (*t.NotificationOverviewData, error) {
	return d.dummy.GetNotificationOverview(ctx, userId)
}

func (d *DataAccessService) GetDashboardNotifications(ctx context.Context, userId uint64, chainIds []uint64, cursor string, colSort t.Sort[enums.NotificationDashboardsColumn], search string, limit uint64) ([]t.NotificationDashboardsTableRow, *t.Paging, error) {
	response := []t.NotificationDashboardsTableRow{}
	var err error

	var currentCursor t.NotificationsDashboardsCursor
	if cursor != "" {
		if currentCursor, err = utils.StringToCursor[t.NotificationsDashboardsCursor](cursor); err != nil {
			return nil, nil, fmt.Errorf("failed to parse passed cursor as NotificationsDashboardsCursor: %w", err)
		}
	}

	// validator query
	vdbQuery := goqu.Dialect("postgres").
		From(goqu.T("users_val_dashboards_notifications_history").As("uvdnh")).
		Select(
			goqu.L("false").As("is_account_dashboard"),
			goqu.I("uvd.network").As("chain_id"),
			goqu.I("uvdnh.epoch"),
			goqu.I("uvd.id").As("dashboard_id"),
			goqu.I("uvd.name").As("dashboard_name"),
			goqu.I("uvdg.id").As("group_id"),
			goqu.I("uvdg.name").As("group_name"),
			goqu.SUM("uvdnh.event_count").As("entity_count"),
			goqu.L("ARRAY_AGG(DISTINCT event_type)").As("event_types"),
		).
		InnerJoin(goqu.T("users_val_dashboards").As("uvd"), goqu.On(
			goqu.Ex{"uvd.id": goqu.I("uvdnh.dashboard_id")})).
		InnerJoin(goqu.T("users_val_dashboards_groups").As("uvdg"), goqu.On(
			goqu.Ex{"uvdg.id": goqu.I("uvdnh.group_id")},
			goqu.Ex{"uvdg.dashboard_id": goqu.I("uvd.id")},
		)).
		Where(
			goqu.Ex{"uvd.user_id": userId},
			goqu.L("uvd.network = ANY(?)", pq.Array(chainIds)),
		).
		GroupBy(
			goqu.I("uvdnh.epoch"),
			goqu.I("uvd.network"),
			goqu.I("uvd.id"),
			goqu.I("uvdg.id"),
			goqu.I("uvdg.name"),
		)

	// TODO account dashboards
	/*adbQuery := goqu.Dialect("postgres").
		From(goqu.T("adb_notifications_history").As("anh")).
		Select(
			goqu.L("true").As("is_account_dashboard"),
			goqu.I("anh.network").As("chain_id"),
			goqu.I("anh.epoch"),
			goqu.I("uad.id").As("dashboard_id"),
			goqu.I("uad.name").As("dashboard_name"),
			goqu.I("uadg.id").As("group_id"),
			goqu.I("uadg.name").As("group_name"),
			goqu.SUM("anh.event_count").As("entity_count"),
			goqu.L("ARRAY_AGG(DISTINCT event_type)").As("event_types"),
		).
		InnerJoin(goqu.T("users_acc_dashboards").As("uad"), goqu.On(
			goqu.Ex{"uad.id": goqu.I("anh.dashboard_id"),
			})).
		InnerJoin(goqu.T("users_acc_dashboards_groups").As("uadg"), goqu.On(
			goqu.Ex{"uadg.id": goqu.I("anh.group_id"),
			goqu.Ex{"uadg.dashboard_id": goqu.I("uad.id")},
			})).
		Where(
			goqu.Ex{"uad.user_id": userId},
			goqu.L("anh.network = ANY(?)", pq.Array(chainIds)),
		).
		GroupBy(
			goqu.I("anh.epoch"),
			goqu.I("anh.network"),
			goqu.I("uad.id"),
			goqu.I("uadg.id"),
			goqu.I("uadg.name"),
		)

	unionQuery := vdbQuery.Union(adbQuery)*/
	unionQuery := goqu.From(vdbQuery)

	// sorting
	defaultColumns := []t.SortColumn{
		{Column: enums.NotificationDashboardTimestamp.ToString(), Desc: true, Offset: currentCursor.Epoch},
		{Column: enums.NotificationDashboardDashboardName.ToString(), Desc: false, Offset: currentCursor.DashboardId},
		{Column: enums.NotificationDashboardGroupName.ToString(), Desc: false, Offset: currentCursor.GroupId},
		{Column: enums.NotificationDashboardChainId.ToString(), Desc: true, Offset: currentCursor.ChainId},
	}
	var offset any
	if currentCursor.IsValid() {
		switch colSort.Column {
		case enums.NotificationDashboardTimestamp:
			offset = currentCursor.Epoch
		case enums.NotificationDashboardDashboardName:
			offset = currentCursor.DashboardId
		case enums.NotificationDashboardGroupName:
			offset = currentCursor.GroupId
		case enums.NotificationDashboardChainId:
			offset = currentCursor.ChainId
		}
	}
	order, directions := applySortAndPagination(defaultColumns, t.SortColumn{Column: colSort.Column.ToString(), Desc: colSort.Desc, Offset: offset}, currentCursor.GenericCursor)
	unionQuery = unionQuery.Order(order...)
	if directions != nil {
		unionQuery = unionQuery.Where(directions)
	}

	// search
	searchName := regexp.MustCompile(`^[a-zA-Z0-9_\-.\ ]+$`).MatchString(search)
	if searchName {
		searchLower := strings.ToLower(strings.Replace(search, "_", "\\_", -1)) + "%"
		unionQuery = unionQuery.Where(exp.NewExpressionList(
			exp.OrType,
			goqu.L("LOWER(?)", goqu.I("dashboard_name")).Like(searchLower),
			goqu.L("LOWER(?)", goqu.I("group_name")).Like(searchLower),
		))
	}
	unionQuery = unionQuery.Limit(uint(limit + 1))

	query, args, err := unionQuery.ToSQL()
	if err != nil {
		return nil, nil, err
	}
	err = d.alloyReader.SelectContext(ctx, &response, query, args...)
	if err != nil {
		return nil, nil, err
	}

	moreDataFlag := len(response) > int(limit)
	if moreDataFlag {
		response = response[:len(response)-1]
	}
	if currentCursor.IsReverse() {
		slices.Reverse(response)
	}
	if !moreDataFlag && !currentCursor.IsValid() {
		// No paging required
		return response, &t.Paging{}, nil
	}
	paging, err := utils.GetPagingFromData(response, currentCursor, moreDataFlag)
	if err != nil {
		return nil, nil, err
	}
	return response, paging, nil
}

func (d *DataAccessService) GetValidatorDashboardNotificationDetails(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, epoch uint64) (*t.NotificationValidatorDashboardDetail, error) {
	return d.dummy.GetValidatorDashboardNotificationDetails(ctx, dashboardId, groupId, epoch)
}

func (d *DataAccessService) GetAccountDashboardNotificationDetails(ctx context.Context, dashboardId uint64, groupId uint64, epoch uint64) (*t.NotificationAccountDashboardDetail, error) {
	return d.dummy.GetAccountDashboardNotificationDetails(ctx, dashboardId, groupId, epoch)
}

func (d *DataAccessService) GetMachineNotifications(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationMachinesColumn], search string, limit uint64) ([]t.NotificationMachinesTableRow, *t.Paging, error) {
	return d.dummy.GetMachineNotifications(ctx, userId, cursor, colSort, search, limit)
}
func (d *DataAccessService) GetClientNotifications(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationClientsColumn], search string, limit uint64) ([]t.NotificationClientsTableRow, *t.Paging, error) {
	return d.dummy.GetClientNotifications(ctx, userId, cursor, colSort, search, limit)
}
func (d *DataAccessService) GetRocketPoolNotifications(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationRocketPoolColumn], search string, limit uint64) ([]t.NotificationRocketPoolTableRow, *t.Paging, error) {
	return d.dummy.GetRocketPoolNotifications(ctx, userId, cursor, colSort, search, limit)
}
func (d *DataAccessService) GetNetworkNotifications(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationNetworksColumn], limit uint64) ([]t.NotificationNetworksTableRow, *t.Paging, error) {
	return d.dummy.GetNetworkNotifications(ctx, userId, cursor, colSort, limit)
}
func (d *DataAccessService) GetNotificationSettings(ctx context.Context, userId uint64) (*t.NotificationSettings, error) {
	return d.dummy.GetNotificationSettings(ctx, userId)
}
func (d *DataAccessService) UpdateNotificationSettingsGeneral(ctx context.Context, userId uint64, settings t.NotificationSettingsGeneral) error {
	return d.dummy.UpdateNotificationSettingsGeneral(ctx, userId, settings)
}
func (d *DataAccessService) UpdateNotificationSettingsNetworks(ctx context.Context, userId uint64, chainId uint64, settings t.NotificationSettingsNetwork) error {
	return d.dummy.UpdateNotificationSettingsNetworks(ctx, userId, chainId, settings)
}
func (d *DataAccessService) UpdateNotificationSettingsPairedDevice(ctx context.Context, userId uint64, pairedDeviceId string, name string, IsNotificationsEnabled bool) error {
	return d.dummy.UpdateNotificationSettingsPairedDevice(ctx, userId, pairedDeviceId, name, IsNotificationsEnabled)
}
func (d *DataAccessService) DeleteNotificationSettingsPairedDevice(ctx context.Context, userId uint64, pairedDeviceId string) error {
	return d.dummy.DeleteNotificationSettingsPairedDevice(ctx, userId, pairedDeviceId)
}
func (d *DataAccessService) UpdateNotificationSettingsClients(ctx context.Context, userId uint64, clientId uint64, IsSubscribed bool) (*t.NotificationSettingsClient, error) {
	return d.dummy.UpdateNotificationSettingsClients(ctx, userId, clientId, IsSubscribed)
}
func (d *DataAccessService) GetNotificationSettingsDashboards(ctx context.Context, userId uint64, cursor string, colSort t.Sort[enums.NotificationSettingsDashboardColumn], search string, limit uint64) ([]t.NotificationSettingsDashboardsTableRow, *t.Paging, error) {
	return d.dummy.GetNotificationSettingsDashboards(ctx, userId, cursor, colSort, search, limit)
}
func (d *DataAccessService) UpdateNotificationSettingsValidatorDashboard(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, settings t.NotificationSettingsValidatorDashboard) error {
	return d.dummy.UpdateNotificationSettingsValidatorDashboard(ctx, dashboardId, groupId, settings)
}
func (d *DataAccessService) UpdateNotificationSettingsAccountDashboard(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, settings t.NotificationSettingsAccountDashboard) error {
	return d.dummy.UpdateNotificationSettingsAccountDashboard(ctx, dashboardId, groupId, settings)
}
