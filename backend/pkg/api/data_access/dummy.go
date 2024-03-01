package dataaccess

import (
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	t "github.com/gobitfly/beaconchain/pkg/api/types"
)

type DummyService struct {
}

func NewDummyService() DummyService {
	return DummyService{}
}

// must pass a pointer to the data
func commonFakeData(a interface{}) error {
	// TODO fake decimal.Decimal
	return faker.FakeData(a, options.WithRandomMapAndSliceMaxSize(5))
}

func (d DummyService) CloseDataAccessService() {
	// nothing to close
}

func (d DummyService) GetUserDashboards(userId uint64) (t.DashboardData, error) {
	r := t.DashboardData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) CreateValidatorDashboard(userId uint64, name string, network uint64) (t.VDBPostReturnData, error) {
	r := t.VDBPostReturnData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardOverview(dashboardId t.VDBIdPrimary) (t.VDBOverviewData, error) {
	r := t.VDBOverviewData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardOverviewByPublicId(publicDashboardId t.VDBIdPublic) (t.VDBOverviewData, error) {
	r := t.VDBOverviewData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardOverviewByValidators(validators t.VDBIdValidatorSet) (t.VDBOverviewData, error) {
	r := t.VDBOverviewData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) RemoveValidatorDashboard(dashboardId t.VDBIdPrimary) error {
	return nil
}

func (d DummyService) RemoveValidatorDashboardByPublicId(dashboardId t.VDBIdPublic) error {
	return nil
}

func (d DummyService) CreateValidatorDashboardGroup(dashboardId t.VDBIdPrimary, name string) (t.VDBOverviewGroup, error) {
	r := t.VDBOverviewGroup{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) CreateValidatorDashboardGroupByPublicId(dashboardId t.VDBIdPublic, name string) (t.VDBOverviewGroup, error) {
	r := t.VDBOverviewGroup{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) RemoveValidatorDashboardGroup(dashboardId t.VDBIdPrimary, groupId uint64) error {
	return nil
}

func (d DummyService) RemoveValidatorDashboardGroupByPublicId(dashboardId t.VDBIdPublic, groupId uint64) error {
	return nil
}

func (d DummyService) AddValidatorDashboardValidators(dashboardId t.VDBIdPrimary, groupId uint64, validators []string) ([]t.VDBPostValidatorsData, error) {
	r := []t.VDBPostValidatorsData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) AddValidatorDashboardValidatorsByPublicId(dashboardId t.VDBIdPublic, groupId uint64, validators []string) ([]t.VDBPostValidatorsData, error) {
	r := []t.VDBPostValidatorsData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardValidators(dashboardId t.VDBIdPrimary, groupId uint64, cursor string, sort []t.Sort[t.VDBManageValidatorsTableColumn], search string, limit uint64) ([]t.VDBManageValidatorsTableRow, t.Paging, error) {
	r := []t.VDBManageValidatorsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardValidatorsByPublicId(dashboardId t.VDBIdPublic, groupId uint64, cursor string, sort []t.Sort[t.VDBManageValidatorsTableColumn], search string, limit uint64) ([]t.VDBManageValidatorsTableRow, t.Paging, error) {
	r := []t.VDBManageValidatorsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardValidatorsByValidators(dashboardId t.VDBIdValidatorSet, cursor string, sort []t.Sort[t.VDBManageValidatorsTableColumn], search string, limit uint64) ([]t.VDBManageValidatorsTableRow, t.Paging, error) {
	r := []t.VDBManageValidatorsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) RemoveValidatorDashboardValidators(dashboardId t.VDBIdPrimary, validators []string) error {
	return nil
}

func (d DummyService) RemoveValidatorDashboardValidatorsByPublicId(dashboardId t.VDBIdPublic, validators []string) error {
	return nil
}

func (d DummyService) CreateValidatorDashboardPublicId(dashboardId t.VDBIdPrimary, name string, showGroupNames bool) (t.VDBPostPublicIdData, error) {
	r := t.VDBPostPublicIdData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) CreateValidatorDashboardPublicIdByPublicId(dashboardId t.VDBIdPublic, name string, showGroupNames bool) (t.VDBPostPublicIdData, error) {
	r := t.VDBPostPublicIdData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) UpdateValidatorDashboardPublicId(dashboardId t.VDBIdPrimary, publicDashboardId string, name string, showGroupNames bool) (t.VDBPostPublicIdData, error) {
	r := t.VDBPostPublicIdData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) UpdateValidatorDashboardPublicIdByPublicId(dashboardId t.VDBIdPublic, publicDashboardId string, name string, showGroupNames bool) (t.VDBPostPublicIdData, error) {
	r := t.VDBPostPublicIdData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) RemoveValidatorDashboardPublicId(dashboardId t.VDBIdPrimary, publicDashboardId string) error {
	return nil
}

func (d DummyService) RemoveValidatorDashboardPublicIdByPublicId(dashboardId t.VDBIdPublic, publicDashboardId string) error {
	return nil
}

func (d DummyService) GetValidatorDashboardSlotViz(dashboardId t.VDBIdPrimary) ([]t.SlotVizEpoch, error) {
	r := []t.SlotVizEpoch{}
	var err error
	for i := 0; i < 4; i++ {
		epoch := t.SlotVizEpoch{}
		err = commonFakeData(&epoch)
		r = append(r, epoch)
	}
	return r, err
}

func (d DummyService) GetValidatorDashboardSlotVizByPublicId(dashboardId t.VDBIdPublic) ([]t.SlotVizEpoch, error) {
	r := []t.SlotVizEpoch{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardSlotVizByValidators(dashboardId t.VDBIdValidatorSet) ([]t.SlotVizEpoch, error) {
	r := []t.SlotVizEpoch{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardSummary(dashboardId t.VDBIdPrimary, cursor string, sort []t.Sort[t.VDBSummaryTableColumn], search string, limit uint64) ([]t.VDBSummaryTableRow, t.Paging, error) {
	r := []t.VDBSummaryTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardSummaryByPublicId(dashboardId t.VDBIdPublic, cursor string, sort []t.Sort[t.VDBSummaryTableColumn], search string, limit uint64) ([]t.VDBSummaryTableRow, t.Paging, error) {
	r := []t.VDBSummaryTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardSummaryByValidators(dashboardId t.VDBIdValidatorSet, cursor string, sort []t.Sort[t.VDBSummaryTableColumn], search string, limit uint64) ([]t.VDBSummaryTableRow, t.Paging, error) {
	r := []t.VDBSummaryTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardGroupSummary(dashboardId t.VDBIdPrimary, groupId uint64) (t.VDBGroupSummaryData, error) {
	r := t.VDBGroupSummaryData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardGroupSummaryByPublicId(dashboardId t.VDBIdPublic, groupId uint64) (t.VDBGroupSummaryData, error) {
	r := t.VDBGroupSummaryData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardGroupSummaryByValidators(dashboardId t.VDBIdValidatorSet) (t.VDBGroupSummaryData, error) {
	r := t.VDBGroupSummaryData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardSummaryChart(dashboardId t.VDBIdPrimary) (t.ChartData[int], error) {
	r := t.ChartData[int]{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardSummaryChartByPublicId(dashboardId t.VDBIdPublic) (t.ChartData[int], error) {
	r := t.ChartData[int]{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardSummaryChartByValidators(dashboardId t.VDBIdValidatorSet) (t.ChartData[int], error) {
	r := t.ChartData[int]{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardRewards(dashboardId t.VDBIdPrimary, cursor string, sort []t.Sort[t.VDBRewardsTableColumn], search string, limit uint64) ([]t.VDBRewardsTableRow, t.Paging, error) {
	r := []t.VDBRewardsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardRewardsByPublicId(dashboardId t.VDBIdPublic, cursor string, sort []t.Sort[t.VDBRewardsTableColumn], search string, limit uint64) ([]t.VDBRewardsTableRow, t.Paging, error) {
	r := []t.VDBRewardsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardRewardsByValidators(dashboardId t.VDBIdValidatorSet, cursor string, sort []t.Sort[t.VDBRewardsTableColumn], search string, limit uint64) ([]t.VDBRewardsTableRow, t.Paging, error) {
	r := []t.VDBRewardsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardGroupRewards(dashboardId t.VDBIdPrimary, groupId uint64, epoch uint64) (t.VDBGroupRewardsData, error) {
	r := t.VDBGroupRewardsData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardGroupRewardsByPublicId(dashboardId t.VDBIdPublic, groupId uint64, epoch uint64) (t.VDBGroupRewardsData, error) {
	r := t.VDBGroupRewardsData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardGroupRewardsByValidators(dashboardId t.VDBIdValidatorSet, epoch uint64) (t.VDBGroupRewardsData, error) {
	r := t.VDBGroupRewardsData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardRewardsChart(dashboardId t.VDBIdPrimary) (t.ChartData[int], error) {
	r := t.ChartData[int]{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardRewardsChartByPublicId(dashboardId t.VDBIdPublic) (t.ChartData[int], error) {
	r := t.ChartData[int]{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardRewardsChartByValidators(dashboardId t.VDBIdValidatorSet) (t.ChartData[int], error) {
	r := t.ChartData[int]{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardDuties(dashboardId t.VDBIdPrimary, epoch uint64, cursor string, sort []t.Sort[t.VDBDutiesTableColumn], search string, limit uint64) ([]t.VDBEpochDutiesTableRow, t.Paging, error) {
	r := []t.VDBEpochDutiesTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardDutiesByPublicId(dashboardId t.VDBIdPublic, epoch uint64, cursor string, sort []t.Sort[t.VDBDutiesTableColumn], search string, limit uint64) ([]t.VDBEpochDutiesTableRow, t.Paging, error) {
	r := []t.VDBEpochDutiesTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardDutiesByValidators(dashboardId t.VDBIdValidatorSet, epoch uint64, cursor string, sort []t.Sort[t.VDBDutiesTableColumn], search string, limit uint64) ([]t.VDBEpochDutiesTableRow, t.Paging, error) {
	r := []t.VDBEpochDutiesTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardBlocks(dashboardId t.VDBIdPrimary, cursor string, sort []t.Sort[t.VDBBlocksTableColumn], search string, limit uint64) ([]t.VDBBlocksTableRow, t.Paging, error) {
	r := []t.VDBBlocksTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardBlocksByPublicId(dashboardId t.VDBIdPublic, cursor string, sort []t.Sort[t.VDBBlocksTableColumn], search string, limit uint64) ([]t.VDBBlocksTableRow, t.Paging, error) {
	r := []t.VDBBlocksTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardBlocksByValidators(dashboardId t.VDBIdValidatorSet, cursor string, sort []t.Sort[t.VDBBlocksTableColumn], search string, limit uint64) ([]t.VDBBlocksTableRow, t.Paging, error) {
	r := []t.VDBBlocksTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardHeatmap(dashboardId t.VDBIdPrimary) (t.VDBHeatmap, error) {
	r := t.VDBHeatmap{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardHeatmapByPublicId(dashboardId t.VDBIdPublic) (t.VDBHeatmap, error) {
	r := t.VDBHeatmap{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardHeatmapByValidators(dashboardId t.VDBIdValidatorSet) (t.VDBHeatmap, error) {
	r := t.VDBHeatmap{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardGroupHeatmap(dashboardId t.VDBIdPrimary, groupId uint64, epoch uint64) (t.VDBHeatmapTooltipData, error) {
	r := t.VDBHeatmapTooltipData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardGroupHeatmapByPublicId(dashboardId t.VDBIdPublic, groupId uint64, epoch uint64) (t.VDBHeatmapTooltipData, error) {
	r := t.VDBHeatmapTooltipData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardGroupHeatmapByValidators(dashboardId t.VDBIdValidatorSet, epoch uint64) (t.VDBHeatmapTooltipData, error) {
	r := t.VDBHeatmapTooltipData{}
	err := commonFakeData(&r)
	return r, err
}

func (d DummyService) GetValidatorDashboardElDeposits(dashboardId t.VDBIdPrimary, cursor string, search string, limit uint64) ([]t.VDBExecutionDepositsTableRow, t.Paging, error) {
	r := []t.VDBExecutionDepositsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardElDepositsByPublicId(dashboardId t.VDBIdPublic, cursor string, search string, limit uint64) ([]t.VDBExecutionDepositsTableRow, t.Paging, error) {
	r := []t.VDBExecutionDepositsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardElDepositsByValidators(dashboardId t.VDBIdValidatorSet, cursor string, search string, limit uint64) ([]t.VDBExecutionDepositsTableRow, t.Paging, error) {
	r := []t.VDBExecutionDepositsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardClDeposits(dashboardId t.VDBIdPrimary, cursor string, search string, limit uint64) ([]t.VDBConsensusDepositsTableRow, t.Paging, error) {
	r := []t.VDBConsensusDepositsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardClDepositsByPublicId(dashboardId t.VDBIdPublic, cursor string, search string, limit uint64) ([]t.VDBConsensusDepositsTableRow, t.Paging, error) {
	r := []t.VDBConsensusDepositsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardClDepositsByValidators(dashboardId t.VDBIdValidatorSet, cursor string, search string, limit uint64) ([]t.VDBConsensusDepositsTableRow, t.Paging, error) {
	r := []t.VDBConsensusDepositsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardWithdrawals(dashboardId t.VDBIdPrimary, cursor string, sort []t.Sort[t.VDBWithdrawalsTableColumn], search string, limit uint64) ([]t.VDBWithdrawalsTableRow, t.Paging, error) {
	r := []t.VDBWithdrawalsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardWithdrawalsByPublicId(dashboardId t.VDBIdPublic, cursor string, sort []t.Sort[t.VDBWithdrawalsTableColumn], search string, limit uint64) ([]t.VDBWithdrawalsTableRow, t.Paging, error) {
	r := []t.VDBWithdrawalsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}

func (d DummyService) GetValidatorDashboardWithdrawalsByValidators(dashboardId t.VDBIdValidatorSet, cursor string, sort []t.Sort[t.VDBWithdrawalsTableColumn], search string, limit uint64) ([]t.VDBWithdrawalsTableRow, t.Paging, error) {
	r := []t.VDBWithdrawalsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, p, err
}
