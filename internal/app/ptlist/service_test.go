package ptlist

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPtListHappyPathAthensTZ(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
	}{
		{
			name:  "Hour test",
			input: []string{"1h", "Europe/Athens", "20210714T204603Z", "20210715T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210714T210000Z",
					"20210714T220000Z",
					"20210714T230000Z",
					"20210715T000000Z",
					"20210715T010000Z",
					"20210715T020000Z",
					"20210715T030000Z",
					"20210715T040000Z",
					"20210715T050000Z",
					"20210715T060000Z",
					"20210715T070000Z",
					"20210715T080000Z",
					"20210715T090000Z",
					"20210715T100000Z",
					"20210715T110000Z",
					"20210715T120000Z"},
			},
		},
		{
			name:  "Day test",
			input: []string{"1d", "Europe/Athens", "20211010T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20211010T210000Z",
					"20211011T210000Z",
					"20211012T210000Z",
					"20211013T210000Z",
					"20211014T210000Z",
					"20211015T210000Z",
					"20211016T210000Z",
					"20211017T210000Z",
					"20211018T210000Z",
					"20211019T210000Z",
					"20211020T210000Z",
					"20211021T210000Z",
					"20211022T210000Z",
					"20211023T210000Z",
					"20211024T210000Z",
					"20211025T210000Z",
					"20211026T210000Z",
					"20211027T210000Z",
					"20211028T210000Z",
					"20211029T210000Z",
					"20211030T210000Z",
					"20211031T220000Z",
					"20211101T220000Z",
					"20211102T220000Z",
					"20211103T220000Z",
					"20211104T220000Z",
					"20211105T220000Z",
					"20211106T220000Z",
					"20211107T220000Z",
					"20211108T220000Z",
					"20211109T220000Z",
					"20211110T220000Z",
					"20211111T220000Z",
					"20211112T220000Z",
					"20211113T220000Z",
					"20211114T220000Z"},
			},
		},
		{
			name:  "Month test",
			input: []string{"1mo", "Europe/Athens", "20210214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210228T220000Z",
					"20210331T210000Z",
					"20210430T210000Z",
					"20210531T210000Z",
					"20210630T210000Z",
					"20210731T210000Z",
					"20210831T210000Z",
					"20210930T210000Z",
					"20211031T220000Z"},
			},
		},
		{
			name:  "Year test",
			input: []string{"1y", "Europe/Athens", "20180214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20181231T220000Z",
					"20191231T220000Z",
					"20201231T220000Z"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.NoError(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}

func TestPtListHappyPathSwedenTZ(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
	}{
		{
			name:  "Hour test",
			input: []string{"1h", "Europe/Stockholm", "20210714T204603Z", "20210715T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210714T210000Z",
					"20210714T220000Z",
					"20210714T230000Z",
					"20210715T000000Z",
					"20210715T010000Z",
					"20210715T020000Z",
					"20210715T030000Z",
					"20210715T040000Z",
					"20210715T050000Z",
					"20210715T060000Z",
					"20210715T070000Z",
					"20210715T080000Z",
					"20210715T090000Z",
					"20210715T100000Z",
					"20210715T110000Z",
					"20210715T120000Z"},
			},
		},
		{
			name:  "Day test",
			input: []string{"1d", "Europe/Stockholm", "20211010T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20211010T210000Z",
					"20211011T210000Z",
					"20211012T210000Z",
					"20211013T210000Z",
					"20211014T210000Z",
					"20211015T210000Z",
					"20211016T210000Z",
					"20211017T210000Z",
					"20211018T210000Z",
					"20211019T210000Z",
					"20211020T210000Z",
					"20211021T210000Z",
					"20211022T210000Z",
					"20211023T210000Z",
					"20211024T210000Z",
					"20211025T210000Z",
					"20211026T210000Z",
					"20211027T210000Z",
					"20211028T210000Z",
					"20211029T210000Z",
					"20211030T210000Z",
					"20211031T220000Z",
					"20211101T220000Z",
					"20211102T220000Z",
					"20211103T220000Z",
					"20211104T220000Z",
					"20211105T220000Z",
					"20211106T220000Z",
					"20211107T220000Z",
					"20211108T220000Z",
					"20211109T220000Z",
					"20211110T220000Z",
					"20211111T220000Z",
					"20211112T220000Z",
					"20211113T220000Z",
					"20211114T220000Z"},
			},
		},
		{
			name:  "Month test",
			input: []string{"1mo", "Europe/Stockholm", "20210214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210228T220000Z",
					"20210331T210000Z",
					"20210430T210000Z",
					"20210531T210000Z",
					"20210630T210000Z",
					"20210731T210000Z",
					"20210831T210000Z",
					"20210930T210000Z",
					"20211031T220000Z"},
			},
		},
		{
			name:  "Year test",
			input: []string{"1y", "Europe/Stockholm", "20180214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20181231T220000Z",
					"20191231T220000Z",
					"20201231T220000Z"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.NoError(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}

func TestPtListHappyPathAbidjanTZ(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
	}{
		{
			name:  "Hour test",
			input: []string{"1h", "Africa/Abidjan", "20210714T204603Z", "20210715T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210714T210000Z",
					"20210714T220000Z",
					"20210714T230000Z",
					"20210715T000000Z",
					"20210715T010000Z",
					"20210715T020000Z",
					"20210715T030000Z",
					"20210715T040000Z",
					"20210715T050000Z",
					"20210715T060000Z",
					"20210715T070000Z",
					"20210715T080000Z",
					"20210715T090000Z",
					"20210715T100000Z",
					"20210715T110000Z",
					"20210715T120000Z"},
			},
		},
		{
			name:  "Day test",
			input: []string{"1d", "Africa/Abidjan", "20211010T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20211010T210000Z",
					"20211011T210000Z",
					"20211012T210000Z",
					"20211013T210000Z",
					"20211014T210000Z",
					"20211015T210000Z",
					"20211016T210000Z",
					"20211017T210000Z",
					"20211018T210000Z",
					"20211019T210000Z",
					"20211020T210000Z",
					"20211021T210000Z",
					"20211022T210000Z",
					"20211023T210000Z",
					"20211024T210000Z",
					"20211025T210000Z",
					"20211026T210000Z",
					"20211027T210000Z",
					"20211028T210000Z",
					"20211029T210000Z",
					"20211030T210000Z",
					"20211031T210000Z",
					"20211101T210000Z",
					"20211102T210000Z",
					"20211103T210000Z",
					"20211104T210000Z",
					"20211105T210000Z",
					"20211106T210000Z",
					"20211107T210000Z",
					"20211108T210000Z",
					"20211109T210000Z",
					"20211110T210000Z",
					"20211111T210000Z",
					"20211112T210000Z",
					"20211113T210000Z",
					"20211114T210000Z"},
			},
		},
		{
			name:  "Month test",
			input: []string{"1mo", "Africa/Abidjan", "20210214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210228T210000Z",
					"20210331T210000Z",
					"20210430T210000Z",
					"20210531T210000Z",
					"20210630T210000Z",
					"20210731T210000Z",
					"20210831T210000Z",
					"20210930T210000Z",
					"20211031T210000Z"},
			},
		},
		{
			name:  "Year test",
			input: []string{"1y", "Africa/Abidjan", "20180214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20181231T210000Z",
					"20191231T210000Z",
					"20201231T210000Z"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.NoError(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}

func TestPtListHappyPathNewYorkTZ(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
	}{
		{
			name:  "Hour test",
			input: []string{"1h", "America/New_York", "20210714T204603Z", "20210715T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210714T210000Z",
					"20210714T220000Z",
					"20210714T230000Z",
					"20210715T000000Z",
					"20210715T010000Z",
					"20210715T020000Z",
					"20210715T030000Z",
					"20210715T040000Z",
					"20210715T050000Z",
					"20210715T060000Z",
					"20210715T070000Z",
					"20210715T080000Z",
					"20210715T090000Z",
					"20210715T100000Z",
					"20210715T110000Z",
					"20210715T120000Z"},
			},
		},
		{
			name:  "Day test",
			input: []string{"1d", "America/New_York", "20211010T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20211010T210000Z",
					"20211011T210000Z",
					"20211012T210000Z",
					"20211013T210000Z",
					"20211014T210000Z",
					"20211015T210000Z",
					"20211016T210000Z",
					"20211017T210000Z",
					"20211018T210000Z",
					"20211019T210000Z",
					"20211020T210000Z",
					"20211021T210000Z",
					"20211022T210000Z",
					"20211023T210000Z",
					"20211024T210000Z",
					"20211025T210000Z",
					"20211026T210000Z",
					"20211027T210000Z",
					"20211028T210000Z",
					"20211029T210000Z",
					"20211030T210000Z",
					"20211031T210000Z",
					"20211101T210000Z",
					"20211102T210000Z",
					"20211103T210000Z",
					"20211104T210000Z",
					"20211105T210000Z",
					"20211106T210000Z",
					"20211107T220000Z",
					"20211108T220000Z",
					"20211109T220000Z",
					"20211110T220000Z",
					"20211111T220000Z",
					"20211112T220000Z",
					"20211113T220000Z",
					"20211114T220000Z"},
			},
		},
		{
			name:  "Month test",
			input: []string{"1mo", "America/New_York", "20210214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210228T220000Z",
					"20210331T210000Z",
					"20210430T210000Z",
					"20210531T210000Z",
					"20210630T210000Z",
					"20210731T210000Z",
					"20210831T210000Z",
					"20210930T210000Z",
					"20211031T210000Z"},
			},
		},
		{
			name:  "Year test",
			input: []string{"1y", "America/New_York", "20180214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20181231T220000Z",
					"20191231T220000Z",
					"20201231T220000Z"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.NoError(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}

func TestPtListHappyPathTokyoTZ(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
	}{
		{
			name:  "Hour test",
			input: []string{"1h", "Asia/Tokyo", "20210214T204603Z", "20210215T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210214T150000Z",
					"20210214T160000Z",
					"20210214T170000Z",
					"20210214T180000Z",
					"20210214T190000Z",
					"20210214T200000Z",
					"20210214T210000Z",
					"20210214T220000Z",
					"20210214T230000Z",
					"20210215T000000Z",
					"20210215T010000Z",
					"20210215T020000Z",
					"20210215T030000Z",
					"20210215T040000Z",
					"20210215T050000Z",
					"20210215T060000Z",
					"20210215T070000Z",
					"20210215T080000Z",
					"20210215T090000Z",
					"20210215T100000Z",
					"20210215T110000Z",
					"20210215T120000Z"},
			},
		},
		{
			name:  "Day test",
			input: []string{"1d", "Asia/Tokyo", "20210214T204603Z", "20210315T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210214T150000Z",
					"20210215T150000Z",
					"20210216T150000Z",
					"20210217T150000Z",
					"20210218T150000Z",
					"20210219T150000Z",
					"20210220T150000Z",
					"20210221T150000Z",
					"20210222T150000Z",
					"20210223T150000Z",
					"20210224T150000Z",
					"20210225T150000Z",
					"20210226T150000Z",
					"20210227T150000Z",
					"20210228T150000Z",
					"20210301T150000Z",
					"20210302T150000Z",
					"20210303T150000Z",
					"20210304T150000Z",
					"20210305T150000Z",
					"20210306T150000Z",
					"20210307T150000Z",
					"20210308T150000Z",
					"20210309T150000Z",
					"20210310T150000Z",
					"20210311T150000Z",
					"20210312T150000Z",
					"20210313T150000Z",
					"20210314T150000Z"},
			},
		},
		{
			name:  "Month test",
			input: []string{"1mo", "Asia/Tokyo", "20210214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210228T150000Z",
					"20210331T150000Z",
					"20210430T150000Z",
					"20210531T150000Z",
					"20210630T150000Z",
					"20210731T150000Z",
					"20210831T150000Z",
					"20210930T150000Z",
					"20211031T150000Z"},
			},
		},
		{
			name:  "Year test",
			input: []string{"1y", "Asia/Tokyo", "20210214T204603Z", "20271115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20211231T150000Z",
					"20221231T150000Z",
					"20231231T150000Z",
					"20241231T150000Z",
					"20251231T150000Z",
					"20261231T150000Z"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.NoError(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}

func TestPtListHappyPathMexicoCityTZ(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
	}{
		{
			name:  "Hour test",
			input: []string{"1h", "America/Mexico_City", "20210214T024603Z", "20210215T053456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210214T060000Z",
					"20210214T070000Z",
					"20210214T080000Z",
					"20210214T090000Z",
					"20210214T100000Z",
					"20210214T110000Z",
					"20210214T120000Z",
					"20210214T130000Z",
					"20210214T140000Z",
					"20210214T150000Z",
					"20210214T160000Z",
					"20210214T170000Z",
					"20210214T180000Z",
					"20210214T190000Z",
					"20210214T200000Z",
					"20210214T210000Z",
					"20210214T220000Z",
					"20210214T230000Z",
					"20210215T000000Z",
					"20210215T010000Z",
					"20210215T020000Z",
					"20210215T030000Z",
					"20210215T040000Z",
					"20210215T050000Z",
					"20210215T060000Z"},
			},
		},
		{
			name:  "Day test",
			input: []string{"1d", "America/Mexico_City", "20210214T024603Z", "20210315T053456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210214T060000Z",
					"20210215T060000Z",
					"20210216T060000Z",
					"20210217T060000Z",
					"20210218T060000Z",
					"20210219T060000Z",
					"20210220T060000Z",
					"20210221T060000Z",
					"20210222T060000Z",
					"20210223T060000Z",
					"20210224T060000Z",
					"20210225T060000Z",
					"20210226T060000Z",
					"20210227T060000Z",
					"20210228T060000Z",
					"20210301T060000Z",
					"20210302T060000Z",
					"20210303T060000Z",
					"20210304T060000Z",
					"20210305T060000Z",
					"20210306T060000Z",
					"20210307T060000Z",
					"20210308T060000Z",
					"20210309T060000Z",
					"20210310T060000Z",
					"20210311T060000Z",
					"20210312T060000Z",
					"20210313T060000Z",
					"20210314T060000Z"},
			},
		},
		{
			name:  "Month test",
			input: []string{"1mo", "America/Mexico_City", "20210214T024603Z", "20211115T053456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210228T060000Z",
					"20210331T060000Z",
					"20210430T050000Z",
					"20210531T050000Z",
					"20210630T050000Z",
					"20210731T050000Z",
					"20210831T050000Z",
					"20210930T050000Z"},
			},
		},
		{
			name:  "Year test",
			input: []string{"1y", "America/Mexico_City", "20210214T024603Z", "20271115T053456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20211231T060000Z",
					"20221231T060000Z",
					"20231231T060000Z",
					"20241231T060000Z",
					"20251231T060000Z",
					"20261231T060000Z"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.NoError(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}

func TestPtListUnhappyPath(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
		expectedError  error
	}{
		{
			name:           "Hour test",
			input:          []string{"1h", "Europe/Aten", "", "20210715T123456Z"},
			expectedOutput: nil,
			expectedError:  errors.New("unknown time zone Europe/Aten"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.Error(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}
